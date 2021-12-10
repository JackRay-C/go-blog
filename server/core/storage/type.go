package storage

import (
	"mime/multipart"
	"path"
	"strings"
	"sync"
)

type FileType int

const (
	IMAGE FileType = iota + 1 // 图片文件
	VIDEO                     // 视频文件
	AUDIO                     // 音频文件
	ZIP                       // 压缩包
	DOCS                      // 文件
	OTHER                     // 其他
)

var fileTypeMap sync.Map

func init() {
	// images
	fileTypeMap.Store(".JPG", IMAGE)
	fileTypeMap.Store(".PNG", IMAGE)
	fileTypeMap.Store(".JPEG", IMAGE)
	fileTypeMap.Store(".GIF", IMAGE)
	fileTypeMap.Store(".BMP", IMAGE)
	fileTypeMap.Store(".ICON", IMAGE)
	fileTypeMap.Store(".PSD", IMAGE)
	// audios
	fileTypeMap.Store(".MP3", AUDIO)
	fileTypeMap.Store(".FLAC", AUDIO)
	fileTypeMap.Store(".WAV", AUDIO)
	fileTypeMap.Store(".MID", AUDIO)

	// videos
	fileTypeMap.Store(".MP4", VIDEO)
	fileTypeMap.Store(".MOV", VIDEO)
	fileTypeMap.Store(".AVI", VIDEO)
	fileTypeMap.Store(".FLV", VIDEO)
	fileTypeMap.Store(".M4V", VIDEO)
	fileTypeMap.Store(".RMVB", VIDEO)

	// docs
	fileTypeMap.Store(".WORD", DOCS)
	fileTypeMap.Store(".DOC", DOCS)
	fileTypeMap.Store(".XLSX", DOCS)
	fileTypeMap.Store(".PPT", DOCS)
	fileTypeMap.Store(".CSV", DOCS)
	fileTypeMap.Store(".XLX", DOCS)
	fileTypeMap.Store(".PDF", DOCS)
	fileTypeMap.Store(".PPTX", DOCS)
	fileTypeMap.Store(".XMIND", DOCS)
	fileTypeMap.Store(".HTML", DOCS)
	fileTypeMap.Store(".JS", DOCS)
	fileTypeMap.Store(".CSS", DOCS)
	fileTypeMap.Store(".GO", DOCS)
	fileTypeMap.Store(".PY", DOCS)
	fileTypeMap.Store(".JAVA", DOCS)
	fileTypeMap.Store(".CLASS", DOCS)
	fileTypeMap.Store(".MD", DOCS)
	fileTypeMap.Store(".YAML", DOCS)
	fileTypeMap.Store(".YML", DOCS)
	fileTypeMap.Store(".INI", DOCS)
	fileTypeMap.Store(".LOG", DOCS)
	fileTypeMap.Store(".JSON", DOCS)
	fileTypeMap.Store(".SH", DOCS)

	// 压缩包
	fileTypeMap.Store(".ZIP", ZIP)
	fileTypeMap.Store(".RAR", ZIP)
	fileTypeMap.Store(".TAR", ZIP)
	fileTypeMap.Store(".TGZ", ZIP)
	fileTypeMap.Store(".TAR.GZ", ZIP)
	fileTypeMap.Store(".TAR.BZ2", ZIP)
	fileTypeMap.Store(".GZ", ZIP)

}


// 通过文件流读取文件头获取文件类型
// todo: getFileTypeByHeader
func GetFileTypeByHeader(header *multipart.FileHeader) FileType {
	panic("doesn't implement this function. ")
}

func GetFileTypeByExt(ext string) FileType  {
	if load, ok := fileTypeMap.Load(strings.ToUpper(ext)); !ok {
		return 0
	} else {
		return load.(FileType)
	}
}

// 通过文件名后缀获取文件类型
func GetFileTypeByFilename(filename string) FileType  {
	ext := strings.ToUpper(path.Ext(filename))
	if load, ok := fileTypeMap.Load(ext); !ok {
		return 0
	} else {
		return load.(FileType)
	}
}

func GetFilePrefix(fileType FileType) string {
	switch int(fileType) {
	case 1:
		return "image"
	case 2:
		return "video"
	case 3:
		return "audio"
	case 4:
		return "zip"
	case 5:
		return "docs"
	default:
		return "other"
	}
}
