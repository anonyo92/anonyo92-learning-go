package iopkg

/* Reading files requires checking most calls for errors. This helper will streamline our error checks below. */
func check(e error) {
	if e != nil {
		panic(e)
	}
}
