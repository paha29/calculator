package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var ip string
	dt := time.Now()
	time := dt.Format("2006-01-02 15:04")
	fmt.Println(time)
	fmt.Print(dt)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pagestring := `	
		<html>
    <head>
        <meta charset="UTF-8">
        <title>Калькулятор блоков</title>
    </head>
    <body>
        <h1>Заполните следующие поля</h1>
        <form method="POST" action="result">
            <label>Длина стенового блока</label><br>
            <input type="radio" name="wBlockLength" value="600" checked> 600<br>
            <input type="radio" name="wBlockLength" value="625" checked> 625<br>
            <label>Ширина стенового блока</label><br>
            <input type="radio" name="wBlockWidth" value="250" checked> 250<br>
            <input type="radio" name="wBlockWidth" value="300" checked> 300<br>
            <input type="radio" name="wBlockWidth" value="375" checked> 350<br>
            <input type="radio" name="wBlockWidth" value="400" checked> 400<br>
            <input type="radio" name="wBlockWidth" value="450" checked> 450<br>
            <input type="radio" name="wBlockWidth" value="500" checked> 500<br>
			<label>Высота стенового блока</label><br>
            <input type="radio" name="wBlockHeight" value="200" checked> 200<br>
            <input type="radio" name="wBlockHeight" value="250" checked> 250<br><br>
            <label>Длина перегородочного блока</label><br>
            <input type="radio" name="pBlockLength" value="600" checked> 600<br>
            <input type="radio" name="pBlockLength" value="625" checked> 625<br>
            <label>Ширина перегородочного блока</label><br>
            <input type="radio" name="pBlockWidth" value="50" checked> 50<br>
            <input type="radio" name="pBlockWidth" value="100" checked> 100<br>
            <input type="radio" name="pBlockWidth" value="150" checked> 150<br>
            <input type="radio" name="pBlockWidth" value="200" checked> 200<br>
            <label>Высота перегородочного блока</label><br>
            <input type="radio" name="pBlockHeight" value="200" checked> 200<br>
            <input type="radio" name="pBlockHeight" value="250" checked> 250<br><br>
            <label>Периметр внешних стен</label><br>
            <input type="text" name="wPerimeter" /><br><br>
            <label>Высота стен дома</label><br>
            <input type="text" name="wHeight" /><br><br>
            <label>Периметр внутренних стен стен</label><br>
            <input type="text" name="pPerimeter" /><br><br>
            <label>Площадь внешних проемов</label><br>
            <input type="text" name="wHole" /><br><br>
            <label>Площадь внутренних проемов</label><br>
            <input type="text" name="pHole" /><br><br>
			<input type="submit" value="Отправить" />
        </form>
    </body>
</html>`
		w.Write([]byte(pagestring))
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		resp, err := http.Get("http://api.sypexgeo.net/xml/")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		for {

			bs := make([]byte, 1024)
			n, err := resp.Body.Read(bs)
			fmt.Println(string(bs[:n]))

			if n == 0 || err != nil {
				break
			}
			f := strings.Replace(string(bs[:n]), `<`, ``, -1)
			k := strings.Split(f, `>`)
			for _, rune := range k[2] {
				if rune >= '0' && rune <= '9' || string(rune) == `.` {
					ip += string(rune)
				}
				if string(rune) == `r` {
					break
				}
			}
			fmt.Print(ip)
		}

		wBlockLength := r.FormValue("wBlockLength")
		wBlockWidth := r.FormValue("wBlockWidth")
		wBlockHeight := r.FormValue("wBlockHeight")
		wlength, err := strconv.ParseFloat(wBlockLength, 64)
		if err != nil {
			fmt.Print(err)
		}
		wwidth, err := strconv.ParseFloat(wBlockWidth, 64)
		if err != nil {
			fmt.Print(err)
		}
		wheight, err := strconv.ParseFloat(wBlockHeight, 64)
		if err != nil {
			fmt.Print(err)
		}

		wBlockVolume := wlength * wwidth * wheight / 1000000000
		wBlockArea := wlength * wheight / 1000000
		fmt.Println(wBlockVolume, wBlockArea)
		wSblock := strconv.FormatFloat(wBlockArea, 'f', 3, 64)
		wVblock := strconv.FormatFloat(wBlockVolume, 'f', 3, 64)

		pBlockLength := r.FormValue("pBlockLength")
		pBlockWidth := r.FormValue("pBlockWidth")
		pBlockHeight := r.FormValue("pBlockHeight")
		plength, err := strconv.ParseFloat(pBlockLength, 64)
		if err != nil {
			fmt.Print(err)
		}
		pwidth, err := strconv.ParseFloat(pBlockWidth, 64)
		if err != nil {
			fmt.Print(err)
		}
		pheight, err := strconv.ParseFloat(pBlockHeight, 64)
		if err != nil {
			fmt.Print(err)
		}

		pBlockVolume := plength * pwidth * pheight / 1000000000
		pBlockArea := plength * pheight / 1000000
		fmt.Println(pBlockVolume, pBlockArea)
		pSblock := strconv.FormatFloat(pBlockArea, 'f', 3, 64)
		pVblock := strconv.FormatFloat(pBlockVolume, 'f', 3, 64)

		wPerimeter := r.FormValue("wPerimeter")
		wPerim := strings.Replace(wPerimeter, ",", ".", -1)
		wPer, err := strconv.ParseFloat(wPerim, 64)
		if err != nil {
			fmt.Print(err)
		}
		wHeight := r.FormValue("wHeight")
		wHeig := strings.Replace(wHeight, ",", ".", -1)
		wHe, err := strconv.ParseFloat(wHeig, 64)
		if err != nil {
			fmt.Print(err)
		}
		pPerimeter := r.FormValue("pPerimeter")
		pPerim := strings.Replace(pPerimeter, ",", ".", -1)
		pPer, err := strconv.ParseFloat(pPerim, 64)
		if err != nil {
			fmt.Print(err)
		}
		wHole := r.FormValue("wHole")
		wHol := strings.Replace(wHole, ",", ".", -1)
		wHo, err := strconv.ParseFloat(wHol, 64)
		if err != nil {
			fmt.Print(err)
		}
		pHole := r.FormValue("pHole")
		pHol := strings.Replace(pHole, ",", ".", -1)
		pHo, err := strconv.ParseFloat(pHol, 64)
		if err != nil {
			fmt.Print(err)
		}

		wArea := wPer*wHe - wHo
		pArea := pPer*wHe - pHo
		wBlockCount := wArea / wBlockArea
		pBlockCount := pArea / pBlockArea
		wBlockCountVolume := wBlockCount * wBlockVolume
		pBlockCountVolume := pBlockCount * pBlockVolume

		wallArea := strconv.FormatFloat(wArea, 'f', 1, 64)
		partitionArea := strconv.FormatFloat(pArea, 'f', 1, 64)
		wallBlockCount := strconv.FormatFloat(wBlockCount, 'f', 0, 64)
		partitionBlockCount := strconv.FormatFloat(pBlockCount, 'f', 0, 64)
		wallBlockCountVolume := strconv.FormatFloat(wBlockCountVolume, 'f', 1, 64)
		partitionBlockCountVolume := strconv.FormatFloat(pBlockCountVolume, 'f', 1, 64)
		pagestring := `	
		<html>
    <head>
        <meta charset="UTF-8">
        <title>Результаты расчета калькулятора</title>
    </head>
    <body>
        <h1>Результаты</h1>
    ` + `Площадь стенового блока = ` + `` + wSblock + ` ` + ` м2` + `<br>
    ` + `Объем стенового блока = ` + `` + wVblock + ` ` + ` м3` + `<br>
    ` + `Площадь перегородочного блока = ` + `` + pSblock + ` ` + ` м2` + `<br>
    ` + `Объем перегородочного блока = ` + `` + pVblock + ` ` + ` м3` + `<br>
    ` + `Площадь внешних стен дома = ` + `` + wallArea + ` ` + ` м2` + `<br>
    ` + `Площадь внутренних стен дома = ` + `` + partitionArea + ` ` + ` м2` + `<br>
    ` + `Необходимо стеновых блоков: ` + `` + wallBlockCount + ` ` + ` шт.` + `<br>
    ` + `Необходимо стеновых блоков: ` + `` + wallBlockCountVolume + ` ` + ` м3` + `<br>
    ` + `Необходимо перегородочных блоков: ` + `` + partitionBlockCount + ` ` + ` шт.` + `<br>
    ` + `Необходимо перегородочных блоков: ` + `` + partitionBlockCountVolume + ` ` + ` м3` + `<br>
    
    </body>
</html>`
		w.Write([]byte(pagestring))

		connStr := "user=postgres password=zxc10asd10qwe20 dbname=calculator_result sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		result, err := db.Exec("insert into calculator (Time, IP, WallValue, PartitionValue) values ($1, $2, $3, $4)", time, ip, wallBlockCountVolume, partitionBlockCountVolume)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected())
	})
	http.ListenAndServe(":8104", nil)
}
