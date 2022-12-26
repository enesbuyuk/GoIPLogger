func GetIPAddress(w http.ResponseWriter, r *http.Request) {
    ip := r.RemoteAddr
    xforward := r.Header.Get("X-Forwarded-For")
    fmt.Println("Url:"+r.URL.Path+"<->IP : "+ip+" <-> X-Forwarded-For : "+ xforward)
}
