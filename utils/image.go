package utils

// import (
// 	"bytes"
// 	"image/jpeg"
// 	"io"
// 	"io/ioutil"

// 	"github.com/nfnt/resize"
// )

// // 调整JPEG图片大小
// // 宽 高
// // inerp算法：NearestNeighbor Bilinear MitchellNetravali Lanczos2 Lanczos3 (0-4)
// func ImgResizeJPEG(width, height uint, r io.Reader,
// 	interp int, w io.Writer) (err error) {
// 	// decode jpeg into image.Image
// 	img, err := jpeg.Decode(r)
// 	if err != nil {
// 		return err
// 	}
// 	m := resize.Resize(width, height, img, resize.InterpolationFunction(interp))

// 	jpeg.Encode(w, m, nil)
// 	return
// }
// func ImgResizeDownJPEG(width, height uint, r io.Reader,
// 	interp int, w io.Writer) (err error) {
// 	// decode jpeg into image.Image
// 	buff, _ := ioutil.ReadAll(r)
// 	img, err := jpeg.Decode(bytes.NewReader(buff))
// 	if err != nil {
// 		return err
// 	}

// 	if uint(img.Bounds().Dx()) < width {
// 		w.Write(buff)
// 		return
// 	}

// 	m := resize.Resize(width, height, img, resize.InterpolationFunction(interp))
// 	jpeg.Encode(w, m, nil)
// 	return
// }
// func ImgThumbnailJPEG(width, height uint, r io.Reader,
// 	interp int, w io.Writer) (err error) {
// 	// decode jpeg into image.Image
// 	img, err := jpeg.Decode(r)
// 	if err != nil {
// 		return err
// 	}
// 	m := resize.Thumbnail(width, height, img, resize.InterpolationFunction(interp))

// 	jpeg.Encode(w, m, nil)
// 	return
// }
