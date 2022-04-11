package main

// Disini teman teman akan mencoba untuk
// melakukan penambahan data pada slice.
// Buatlah variable slice dengan tipe data string.
// Lalu masukan nama kalian ke dalam slice.
// Expected outout: ["NamaPanggilan", "Nama Akhir"]
// Contoh [Zein Fahrozi]
// Outputkan jawabannya ya pastikan cap dan len nya adalah 2
func main() {
	// TODO: answer here
}

// AddToSlice is function to add string parameter to slice data
func AddToSlice(fn string, ln string) []string {
	var sliceData []string

	sliceData = append(sliceData, fn)
	sliceData = append(sliceData, ln)

	return sliceData
}
