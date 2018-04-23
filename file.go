package main

import (
	"os"
	"fmt"
)

func DeepDir(dir *os.File, deep bool, symlink bool) (rtn []os.FileInfo) {
	info, err := dir.Readdir(0)
	//if rtn == nil {
	//	rtn = make([]os.FileInfo, 0, 20)
	//}
	if err != nil {
		return rtn
	}
	for _, f := range info {
		if f.Name()[0] == '.' ||f.Name()[0] == '_' {
			continue
		}
		if !f.IsDir() {
			if symlink && f.Mode() & os.ModeSymlink != 0 {
				//符号链接文件也列出
				rtn = append(rtn, f)
			} else {
				rtn = append(rtn, f)
			}
			continue
		}
		if deep {
			subdir, err := os.Open(dir.Name()+string(os.PathSeparator)+f.Name())
			defer subdir.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr,err)
			}
			rtn = append(rtn, DeepDir(subdir, deep, symlink)...)
			continue
		}
	}
	return rtn
}
func Notempty(in []os.FileInfo) []os.FileInfo {
	i := 0
	for _, e := range in {
		if e != nil {
			in[i] = e
			i++
		}
	}
	return in[:i]
}

func main() {
	dir, err := os.Open("C:\\Users\\modian\\Downloads")
	defer dir.Close()
	if err != nil {
		panic(err)
	}
	for _, f := range DeepDir(dir, true, true) {
		fmt.Println(f.Name(), f.IsDir(), f.Mode() & os.ModeDir != 0)
	}

}
