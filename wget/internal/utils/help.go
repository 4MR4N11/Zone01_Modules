package utils

import "fmt"

func GetHelp() {
	fmt.Println("Wget, a non-interactive network retriever.")
	fmt.Println("Usage: wget [OPTION]... [URL]...")
	fmt.Println()
	fmt.Println("Available options:")
	fmt.Println()
	fmt.Println("  Download:")
	fmt.Println("   -B                            download a file(URL) immediately to the background")
	fmt.Println("                                 and the output should be redirected to a log file(wget-log)")
	fmt.Println("   -O=FILE                       download a file and save it under a different name")
	fmt.Println("   -P=DIRECTORY/                 save files to DIRECTORY/")
	fmt.Println("   --rate-limit=LIMIT            limits the speed of download to LIMIT")
	fmt.Println("   -i=FILE                       download URLs found in local or external FILE")
	fmt.Println()
	fmt.Println("  Mirror:")
	fmt.Println("   --mirror                      retrieve(recursively) and parse the HTML or CSS from the given URL")
	fmt.Println("   optional flags to go along with the --mirror:")
	fmt.Println("      -R,  --reject=SUFFIXES     avoid downloading list of file SUFFIXES")
	fmt.Println("      -X,  --exclude=PATHS       avoid downloading list of PATHS")
	fmt.Println("      --convert-links            make links in downloaded HTML or CSS point to local files")
}
