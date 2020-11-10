package response

import "os"

func ExportFile(output string) (int, string) {
	buff := []byte(output)
	f, err := os.Create("./export/output.txt")
	if err != nil {
		return 500, "error creating document"
	}
	_, err = f.Write(buff)
	if err != nil {
		return 500, "error writing to document"
	}
	f.Sync()
	return 200, "OK"
}
