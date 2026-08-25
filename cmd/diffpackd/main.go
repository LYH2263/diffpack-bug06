package main

import (
    "encoding/json"
    "flag"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/LYH2263/go-diffpack"
)

func main() {
    addr := flag.String("addr", ":8240", "listen")
    node := flag.String("node", "local", "node id")
    flag.Parse()

    p, err := diffpack.New(diffpack.Options{NodeID: *node, AuditPath: "audit"})
    if err != nil {
        log.Fatal(err)
    }
    defer p.Close()

    web := "web"
    if _, err := os.Stat(web); err != nil {
        web = filepath.Join("..", "..", "web")
    }

    mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(web))))
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        http.ServeFile(w, r, filepath.Join(web, "index.html"))
    })
    mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, p.Health())
    })
    mux.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, p.Stats())
    })
    mux.HandleFunc("/api/jobs", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            writeJSON(w, p.ListJobs())
            return
        }
        if r.Method != http.MethodPost {
            http.Error(w, "method", http.StatusMethodNotAllowed)
            return
        }
        body, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        var req struct {
            ID     string `json:"id"`
            Base   []byte `json:"base"`
            Target []byte `json:"target"`
        }
        if err := json.Unmarshal(body, &req); err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        _, err = p.BuildDelta(r.Context(), req.ID, req.Base, req.Target)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        job, _ := p.GetJob(req.ID)
        writeJSON(w, job)
    })
    mux.HandleFunc("/api/jobs/", func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Path[len("/api/jobs/"):]
        if id == "" {
            http.NotFound(w, r)
            return
        }
        if r.URL.Query().Get("action") == "hunks" {
            hunks, err := p.ListHunks(id)
            if err != nil {
                http.Error(w, err.Error(), 404)
                return
            }
            writeJSON(w, hunks)
            return
        }
        if r.URL.Query().Get("action") == "verify" {
            recJob, err := p.GetJob(id)
            if err != nil {
                http.Error(w, err.Error(), 404)
                return
            }
            _ = recJob
            if err := p.MarkVerified(id); err != nil {
                http.Error(w, err.Error(), 500)
                return
            }
            writeJSON(w, map[string]string{"status": "verified"})
            return
        }
        job, err := p.GetJob(id)
        if err != nil {
            http.Error(w, err.Error(), 404)
            return
        }
        writeJSON(w, job)
    })

    log.Printf("diffpackd %s", *addr)
    log.Fatal(http.ListenAndServe(*addr, mux))
}

func writeJSON(w http.ResponseWriter, v interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}
