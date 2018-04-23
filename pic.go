package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"sync"
	"google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

//读取图片时间
func ReadPicTime(file string) (string, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return file, errors.New("图片文件不存在")
	}
	fmt.Printf("文件名:%s\n", file)
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(bytes.NewReader(buf))
	if err != nil {
		return file, errors.New("图片扩展信息无法读取")
	}

	if tm, e := x.DateTime(); e == nil {
		return fmt.Sprintf("%4d%02d%02d_%02d%02d%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second()), nil
	}
	return file, errors.New("图片时间信息无法读取")
}

//判断图片时间，判断是否已存在文件，重命名
func RenamePic(dir, src, dest string, renameit bool) (string, error) {
	if err := os.Chdir(dir); err != nil {
		fmt.Println("目录不存在:" + dir)
		return src, errors.New("目录不存在:" + dir)
	}
	if f, err := os.Open(dest); err == nil {
		f.Close()
		fmt.Println("\t目标文件已存在", dest, "尝试新名字")
		var seq int64 = 1
		suffix := strings.ToUpper(dest[strings.LastIndex(dest, ".")+1:])
		base := dest[:strings.LastIndex(dest, ".")]
		if strings.Count(base, "_") > 1 {
			seq, _ = strconv.ParseInt(dest[strings.LastIndex(dest, "_")+1:strings.LastIndex(dest, ".")], 10, 32)
			base = base[:strings.LastIndex(dest, "_")]
			seq++
		}
		return RenamePic(dir, src, fmt.Sprintf("%s_%02d.%s", base, seq, suffix), renameit)
	}

	if f, err := os.Open(src); err != nil && os.IsNotExist(err) {
		fmt.Println("原文件不存在:" + dir)
		return dest, errors.New("原文件不存在:" + src)
	} else {
		f.Close()
	}
	if renameit {
		if err := os.Rename(src, dest); err != nil {
			return dest, err
		}
	}
	return dest, nil
}

//处理图片文件，根据名称和时间进行重命名
func Process(f string) {
	if tm, err := ReadPicTime(f); err == nil {
		fmt.Println("\t", f, "日期", tm)
		if strings.HasPrefix(f, tm) {
			//命名符合规则，忽略
			fmt.Println("\t忽略图片", f)
		} else {
			dest, err := RenamePic(".", f, fmt.Sprintf("%s.%s", tm, strings.ToUpper(f[strings.LastIndex(f, ".")+1:])), true)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("\tRename ok:", f, dest)
			}
		}
	} else {
		fmt.Println("\t", f, "日期无法读取")
	}
	fmt.Println()
}

func main() {
	fm := map[string]string{
		"JPG": "JPG", "JPEG": "JPEG",
		"NEF": "NEF", "ARW": "ARW", "CR2": "CR2",
	}
	defer func() {
		fmt.Println("运行结束")
	}()
	if len(os.Args) > 1 {
		for _, f := range os.Args[1:] {
			if _, ok := fm[strings.ToUpper(f[strings.LastIndex(f, ".")+1:])]; ok {
				//ReadPic(f)
				Process(f)
			} else {
				fmt.Println("忽略非图文件：", f)
			}
		}
	} else {
		//读当前目录
		info, _ := ioutil.ReadDir(".")
		//var files[] string
		files := make([]string, 1, 10240)
		for _, fi := range info {
			if fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			//if strings.HasSuffix(strings.ToUpper(f.Name()), "ARW") || strings.HasSuffix(strings.ToUpper(f.Name()), "JPG") || strings.HasSuffix(strings.ToUpper(f.Name()), "JPEG") || strings.HasSuffix(strings.ToUpper(f.Name()), "NEF") {
			f := fi.Name()
			if _, ok := fm[strings.ToUpper(f[strings.LastIndex(f, ".")+1:])]; ok {
				//Process(f) // 单coroutine顺序处理
				files = append(files, f)
			} else {
				fmt.Println("忽略非图文件：", f)
			}
		}
		done := make(chan bool, 1)
		fmt.Printf("图片数:%d\n", len(files))
		//// 并发处理
		wg := sync.WaitGroup{}
		//每个图片一个协程
		for _, f := range files {
			wg.Add(1)
			go func(f string) {
				defer wg.Done()
				Process(f)
			}(f)
		}
		wg.Wait()

		done <- true

		select {
		case <-done:
			fmt.Printf("done")
		}
		fmt.Println("图片处理结束")
	}
}

//读取图片文件exif信息，支持JPG、NEF、ARW，只有JPG格式才读取宽高信息
func ReadPic(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("文件名:%s\n", file)
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(bytes.NewReader(buf))
	if err != nil {
		return
	}
	if v, e := x.Get(exif.ExposureTime); e == nil {
		fmt.Printf("曝光时间:%s\n", v)
	}
	if v, e := x.Get(exif.FocalLengthIn35mmFilm); e == nil {
		fmt.Printf("35mm焦距:%s\n", v)
	}
	if v, e := x.Get(exif.FNumber); e == nil {
		fmt.Printf("光圈:%s\n", v)
	}
	if v, e := x.Get(exif.ISOSpeedRatings); e == nil {
		fmt.Printf("ISO:%s\n", v)
	}
	if v, e := x.Get(exif.XResolution); e == nil {
		fmt.Printf("水平分辨率%sdpi\n", v)
	}
	if v, e := x.Get(exif.YResolution); e == nil {
		fmt.Printf("垂直分辨率%sdpi\n", v)
	}
	if v, e := x.Get(exif.PixelYDimension); e == nil {
		fmt.Println("PixelYDimension:", v)
	}
	if v, e := x.Get(exif.ResolutionUnit); e == nil {
		fmt.Println("分辨率率单位:", v)
	}
	if v, e := x.Get(exif.ExposureProgram); e == nil {
		fmt.Println("曝光方式:", v)
	}
	if v, e := x.Get(exif.MaxApertureValue); e == nil {
		fmt.Println("最大光圈:", v)
	}
	if v, e := x.Get(exif.WhiteBalance); e == nil {
		fmt.Println("白平衡:", v)
	}
	if v, e := x.Get(exif.LensMake); e == nil {
		fmt.Println("镜头:", v)
	}
	if v, e := x.Get(exif.LensModel); e == nil {
		fmt.Println("镜头参数:", v)
	}
	if v, e := x.Get(exif.ImageWidth); e == nil {
		fmt.Println("ImageWidth:", v)
	}
	if v, e := x.Get(exif.Model); e == nil {
		fmt.Println("相机:", v)
	}
	if v, e := x.Get(exif.Make); e == nil {
		fmt.Println("制造商:", v)
	}
	if v, e := x.Get(exif.Software); e == nil {
		fmt.Println("软件:", v)
	}
	if v, e := x.Get(exif.Copyright); e == nil {
		fmt.Println("版权:", v)
	}
	if tm, e := x.DateTime(); e == nil {
		fmt.Printf("时间:%4d%02d%02d_%02d%02d%02d\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	} else if v, e := x.Get(exif.DateTimeOriginal); e == nil {
		fmt.Println("拍摄时间:", strings.TrimSpace(v.String()))
	} else if v, e := x.Get(exif.DateTimeDigitized); e == nil {
		fmt.Println("数字化时间:", strings.TrimSpace(string(v.Val)))
	} else {
		fmt.Println("时间:", file)
	}
	fmt.Println("文件大小:", len(buf))
	if strings.HasSuffix(strings.ToUpper(file), "JPG") || strings.HasSuffix(strings.ToUpper(file), "JPEG") {
		conf, format, err := image.DecodeConfig(bytes.NewReader(buf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\t%s", file, err.Error())
			return
		}
		fmt.Println(format, conf.Width, "x", conf.Height)
	}
}
