func GetIPAddress(w http.ResponseWriter, r *http.Request) {
    ip := r.RemoteAddr
    xforward := r.Header.Get("X-Forwarded-For")
    fmt.Println("Url:"+r.URL.Path+"<->IP : "+ip+" <-> X-Forwarded-For : "+ xforward)
}
func AppendToFile(w http.ResponseWriter, r *http.Request, message string) error {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	fileName := "main.log"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Yazma işlemini gerçekleştir
	writer := bufio.NewWriter(file)

	// Mesajı dosyaya yaz
	now := time.Now()
	timeString := now.Format("2006-01-02 15:04:05")
	_, err = writer.WriteString(timeString + " | " + message + "--->Url:" + r.URL.Path + "<->IP : " + ip + " <-> X-Forwarded-For : " + xforward + "\n")
	if err != nil {
		return err
	}
	// Değişiklikleri dosyaya kaydet
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
