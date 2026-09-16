package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const manifestName = ".mrpack-files"

type index struct {
	Dependencies map[string]string `json:"dependencies"`
	Files        []struct {
		Path      string            `json:"path"`
		Hashes    map[string]string `json:"hashes"`
		Env       map[string]string `json:"env"`
		Downloads []string          `json:"downloads"`
	} `json:"files"`
}

type project struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	if len(os.Args) != 3 {
		panic("usage: import_mrpack <pack.mrpack> <destination>")
	}

	destination, err := filepath.Abs(os.Args[2])
	must(err)
	removePreviousImport(destination)

	r, err := zip.OpenReader(os.Args[1])
	must(err)
	defer r.Close()

	var idx index
	for _, f := range r.File {
		if f.Name == "modrinth.index.json" {
			rc, err := f.Open()
			must(err)
			must(json.NewDecoder(rc).Decode(&idx))
			must(rc.Close())
			break
		}
	}
	if len(idx.Files) == 0 {
		panic("modrinth.index.json is missing or contains no files")
	}
	updatePackVersions(destination, idx.Dependencies)

	imported := make([]string, 0, len(idx.Files))
	titles := loadProjectTitles(idx)
	for _, f := range idx.Files {
		if len(f.Downloads) != 1 {
			panic(fmt.Sprintf("%s has %d download URLs", f.Path, len(f.Downloads)))
		}

		projectID, versionID := modrinthIDs(f.Downloads[0])
		name := titles[projectID]
		if name == "" {
			name = strings.TrimSuffix(filepath.Base(f.Path), filepath.Ext(f.Path))
		}

		side := "both"
		if f.Env["server"] == "unsupported" {
			side = "client"
		} else if f.Env["client"] == "unsupported" {
			side = "server"
		}

		hashFormat := "sha1"
		hash := f.Hashes[hashFormat]
		if hash == "" {
			hashFormat = "sha512"
			hash = f.Hashes[hashFormat]
		}

		metaPath := strings.TrimSuffix(f.Path, filepath.Ext(f.Path)) + ".pw.toml"
		must(os.MkdirAll(filepath.Dir(filepath.Join(destination, metaPath)), 0o755))
		content := fmt.Sprintf("name = %s\nfilename = %s\nside = %s\n\n[download]\nhash-format = %s\nhash = %s\nurl = %s\n",
			strconv.Quote(name), strconv.Quote(filepath.Base(f.Path)), strconv.Quote(side),
			strconv.Quote(hashFormat), strconv.Quote(hash), strconv.Quote(f.Downloads[0]))
		if projectID != "" && versionID != "" {
			content += fmt.Sprintf("\n[update.modrinth]\nmod-id = %s\nversion = %s\n",
				strconv.Quote(projectID), strconv.Quote(versionID))
		}
		must(os.WriteFile(filepath.Join(destination, metaPath), []byte(content), 0o644))
		imported = append(imported, filepath.ToSlash(metaPath))
	}

	for _, f := range r.File {
		const prefix = "overrides/"
		if !strings.HasPrefix(f.Name, prefix) || f.FileInfo().IsDir() {
			continue
		}
		rel := strings.TrimPrefix(f.Name, prefix)
		dst := safePath(destination, rel)
		must(os.MkdirAll(filepath.Dir(dst), 0o755))
		rc, err := f.Open()
		must(err)
		out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		must(err)
		_, err = io.Copy(out, rc)
		must(err)
		must(out.Close())
		must(rc.Close())
		imported = append(imported, filepath.ToSlash(rel))
	}

	sort.Strings(imported)
	must(os.WriteFile(filepath.Join(destination, manifestName), []byte(strings.Join(imported, "\n")+"\n"), 0o644))
}

func updatePackVersions(destination string, versions map[string]string) {
	path := filepath.Join(destination, "pack.toml")
	data, err := os.ReadFile(path)
	must(err)
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		key := strings.TrimSpace(strings.SplitN(line, "=", 2)[0])
		if version, ok := versions[key]; ok {
			lines[i] = key + " = " + strconv.Quote(version)
		}
	}
	must(os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644))
}

func removePreviousImport(destination string) {
	f, err := os.Open(filepath.Join(destination, manifestName))
	if os.IsNotExist(err) {
		return
	}
	must(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			must(os.Remove(safePath(destination, scanner.Text())))
		}
	}
	must(scanner.Err())
}

func safePath(root, relative string) string {
	dst := filepath.Join(root, filepath.FromSlash(relative))
	cleanDst, err := filepath.Abs(dst)
	must(err)
	if !strings.HasPrefix(cleanDst, root+string(os.PathSeparator)) {
		panic("unsafe import path: " + relative)
	}
	return cleanDst
}

func loadProjectTitles(idx index) map[string]string {
	seen := map[string]bool{}
	ids := make([]string, 0, len(idx.Files))
	for _, f := range idx.Files {
		if len(f.Downloads) != 1 {
			continue
		}
		id, _ := modrinthIDs(f.Downloads[0])
		if id != "" && !seen[id] {
			seen[id] = true
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	b, err := json.Marshal(ids)
	must(err)
	resp, err := http.Get("https://api.modrinth.com/v2/projects?ids=" + url.QueryEscape(string(b)))
	must(err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Modrinth API returned " + resp.Status)
	}
	var projects []project
	must(json.NewDecoder(resp.Body).Decode(&projects))
	titles := make(map[string]string, len(projects))
	for _, p := range projects {
		titles[p.ID] = p.Title
	}
	return titles
}

func modrinthIDs(raw string) (string, string) {
	u, err := url.Parse(raw)
	if err != nil || u.Host != "cdn.modrinth.com" {
		return "", ""
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 5 || parts[0] != "data" || parts[2] != "versions" {
		return "", ""
	}
	return parts[1], parts[3]
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
