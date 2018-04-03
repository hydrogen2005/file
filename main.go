package main
/**
遍历path下的所有文件和文件夹
 */
import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"io"
	//"strconv"
	"strings"
)
type MyFile struct{
	name string
	dir string
	size int64
}

func main(){
	//读取2个文件夹下所有的文件
	pathA:="G:\\5"
	pathB:="G:\\6"
	//pachC:="G:\\3"
	pachD:="G:\\4"
	difFiles:=make(map[int64] []MyFile)
	filesA:=getFiles(pathA)
	filesB:=getFiles(pathB)
	//找到大小相同的文件，输出他们的path
	for _,file:=range filesA{
		size:=file.size
		files:=difFiles[size]
		if(files!=nil){//已经有数据了
			difFiles[size]=append(files,file)
		}else{
			difFiles[size]=append(make([]MyFile,0),file)
		}
	}
	for _,file:=range filesB{
		size:=file.size
		files:=difFiles[size]
		if(files!=nil){//已经有数据了
			difFiles[size]=append(files,file)
		}else{
			difFiles[size]=append(make([]MyFile,0),file)
		}
	}

	//重命名相同的文件为indexA ，indexB,并输出到一个临时目录中
	//i:=0
	// for size,files:=range difFiles{
	// 	if(len(files)>1){
	// 		fmt.Println("dif size=",size," files=",files)
	// 		j:=0
	//		for _,file:=range files{
	//			source:=file.dir+"\\"+file.name
	//			dest:=pachC+"\\"+strconv.Itoa(i)+"\\"+strconv.Itoa(j)+file.name
	//			fmt.Println("source=",source,"|dest=",dest)
	//			//创建dest目录
	//			err:=os.MkdirAll(pachC+"\\"+strconv.Itoa(i),os.ModeDir|os.ModePerm)
	//			if(err!=nil){
	//				fmt.Println(err)
	//				return
	//			}
	//			copyFile(source,dest)
	//			j++
	//		}
	//		i++
	//	}
	// }
	//人工检测是否文件是否相同
	//移动文件数组中的第一个到指定文件夹
	for size,files:=range difFiles {
			fmt.Println("dif size=", size, " files=", files)
			for _, file := range files {
				source:=file.dir+"\\"+file.name
				dest:=pachD+"\\"+file.name
				fmt.Println("copy file source=",source,"dest=",dest)
				copyFile(source,dest)
				break
			}
	}

}
//获得指定path下的全部文件
func getFiles(path string) []MyFile{
	//files:= make(map[string]int64)
	//file, _ := os.Getwd()
	//fmt.Println(file)
	//path := file
	files:=make([]MyFile,0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//fmt.Println(info.Size())
		if info.IsDir() {

		} else {
			ss:=strings.Split(path,"\\")
			//fmt.Println(ss[len(ss)-1])
			dir:=ss[0:(len(ss)-1)]
			dir2:=strings.Join(dir,"\\")
			fmt.Println(dir2)
			name:=info.Name()
			fmt.Println(name)
			size :=info.Size()

			var file=MyFile{
				name,
				dir2,
				size,
			}
			fmt.Println(file)
			files=append(files,file)
		}
		return nil
	})

	return files
}

//拷贝文件  要拷贝的文件路径 拷贝到哪里
func copyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}
	//打开文件资源
	source_open, err := os.Open(source)
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()
	//只写模式打开文件 如果文件不存在进行创建 并赋予 644的权限。详情查看linux 权限解释
	dest_open, err:= os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		log.Println("openFile error",err.Error())
		return false
	}
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}
