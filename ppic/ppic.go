package ppic

import (
	"crypto/md5"
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"bytes"
	"strconv"
	"errors"
	"strings"
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

var fm = map[string]string{
"JPG": "JPG", "JPEG": "JPEG",
"NEF": "NEF", "ARW": "ARW", "CR2": "CR2",
}

var sema = make(chan struct{}, 20)

func RenameAllPic(root string) ([]string, error) {
	files := make([]string,10)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}
		file := info.Name()
		// is pic.
		if _, ok := fm[strings.ToUpper(file[strings.LastIndex(file, ".")+1:])]; ok {

		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
