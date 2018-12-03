// 函数

func protect(g func()){
	defer func(){
		log.PrintLn("done")
		if x:= recover(),x !=nil {
			log.Printf("run time panic: %v", x)
		}
	}{}
	log.Println("start")
	g()
}