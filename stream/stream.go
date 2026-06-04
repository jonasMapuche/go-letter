package stream

const (
	ONE = 0
)

/*
func Video() *gocv.VideoCapture {
	var webcam *gocv.VideoCapture
	var err error
	deviceID := ONE
	webcam, err = gocv.VideoCaptureDevice(deviceID)
	checkErr(err)
	if !webcam.IsOpened() {
		return webcam
	}
	return webcam
}

func Read(webcam *gocv.VideoCapture, image gocv.Mat) *gocv.NativeByteBuffer {
	if ok := webcam.Read(&image); !ok || image.Empty() {
		return nil
	}
	var buffer *gocv.NativeByteBuffer
	buffer, err := gocv.IMEncode(".jpg", image)
	checkErr(err)
	return buffer
}
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
