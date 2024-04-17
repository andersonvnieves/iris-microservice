Caminho do gerador de certificado

C:\Program Files\Go\src\crypto\tls\generate_cert.go



comando go run 'C:\Program Files\Go\src\crypto\tls\generate_cert.go' --host localhost




add this to launch.json to allow interations on integrated VS Code terminal
 "console": "integratedTerminal"





//Sample add headers 
type addIrisHandler string

func (addIris addIrisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Tina", "Lotus")
	http.SetCookie(w, &http.Cookie{
		Name:    "biscoiteira",
		Value:   "0923984",
		Expires: time.Now().Add(24 * time.Hour),
	})
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(mh))
}


//Serve File



	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/servefile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./iris.csv")
	})

	http.HandleFunc("/iris", func(w http.ResponseWriter, r *http.Request) {
		irisFile, err := os.Open("./iris.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer irisFile.Close()

		data, err := io.ReadAll(irisFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(data))
	})