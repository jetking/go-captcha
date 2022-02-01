/**
 * @Au.charshor Awen
 * @Description
 * @Date 2021/7/20
 **/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/wenlng/go-captcha/captcha"
)

// go test -race base.go captcha_test.go
func TestGetCaptchaGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	n := 30
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			t.Logf(">>> %p\n", getCaptcha())
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestGenerate(t *testing.T) {
	capt := captcha.GetCaptcha()
	capt.SetTextRangLen(captcha.RangeVal{Min: 5, Max: 7})
	dot, data, tData, key, err := capt.Generate()
	if err != nil {
		t.Log("err:", err)
		t.FailNow()
	}
	fmt.Printf("dot:%#v\n", dot)
	fmt.Println("key:", key)

	filename := captcha.RandInt(1, 200)
	file := getPWD() + "/tests/.cache/" + fmt.Sprintf("%v", filename) + ".jpg"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("os.OpenFile() fail:", err)
		t.FailNow()
	}
	defer f.Close()
	dataReader := bytes.NewReader(data)
	if _, err = io.Copy(f, dataReader); err != nil {
		fmt.Println("io.Copy() fail:", err)
	}

	file = getPWD() + "/tests/.cache/" + fmt.Sprintf("%v", filename) + "_t.jpg"
	f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("os.OpenFile() fail:", err)
		t.FailNow()
	}
	defer f.Close()
	dataReader = bytes.NewReader(tData)
	if _, err = io.Copy(f, dataReader); err != nil {
		fmt.Println("io.Copy() fail:", err)
	}

}

func TestImageSize(t *testing.T) {
	capt := getCaptcha()
	fmt.Println(capt)
	capt.SetImageSize(captcha.Size{Width: 300, Height: 300})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestSetThumbSize(t *testing.T) {
	capt := getCaptcha()

	capt.SetThumbSize(captcha.Size{Width: 300, Height: 300})

	chars := []string{"HE", "CA", "WO", "NE", "HT", "IE", "PG", "GI", "CH", "CO", "DA"}
	_ = capt.SetRangChars(chars)
	//capt.SetImageFontDistort(0)
	//capt.SetImageFontDistort(0)
	dots, data, tData, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	file := getPWD() + "/tests/.cache/" + fmt.Sprintf("%v", captcha.RandInt(1, 200)) + "Img.png"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	dataReader := bytes.NewReader(data)
	_, err = io.Copy(logFile, dataReader)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(data))
	fmt.Println(len(tData))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestChars(t *testing.T) {
	capt := getCaptcha()
	//chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//capt.SetRangChars(strings.Split(chars, ""))
	chars := []string{"HE", "CA", "WO", "NE", "HT", "IE", "PG", "GI", "CH", "CO", "DA"}
	_ = capt.SetRangChars(chars)
	//chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
	//capt.SetRangChars(chars)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestColors(t *testing.T) {
	capt := getCaptcha()
	capt.SetTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
		"#712217",
		"#885500",
		"#392585",
	})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestAlpha(t *testing.T) {
	capt := getCaptcha()

	capt.SetImageFontAlpha(0.5)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestImageFontDistort(t *testing.T) {
	capt := getCaptcha()

	capt.SetImageFontDistort(captcha.DistortLevel2)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestRangAnglePos(t *testing.T) {
	capt := getCaptcha()

	rang := []captcha.RangeVal{
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
	}
	capt.SetTextRangAnglePos(rang)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBackground(t *testing.T) {
	capt := getCaptcha()

	//capt.SetThumbBackground([]string{
	//	getPWD() + "/__example/resources/images/thumb/r1.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r2.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r3.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r4.jpg",
	//	getPWD() + "/__example/resources/images/thumb/r5.jpg",
	//})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBgCircles(t *testing.T) {
	capt := getCaptcha()

	capt.SetThumbBgCirclesNum(200)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}
