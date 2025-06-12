package main

import (
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go File Operations Examples ===")

	// 1. Basic file creation and writing
	fmt.Println("1. Basic file operations:")
	basicFileOperations()

	// 2. Reading files different ways
	fmt.Println("\n2. Reading files:")
	readFileExamples()

	// 3. File info and stats
	fmt.Println("\n3. File information:")
	fileInfoExamples()

	// 4. Directory operations
	fmt.Println("\n4. Directory operations:")
	directoryOperations()

	// 5. File permissions and modes
	fmt.Println("\n5. File permissions:")
	filePermissions()

	// 6. Temporary files and directories
	fmt.Println("\n6. Temporary files:")
	tempFileOperations()

	// 7. File copying and moving
	fmt.Println("\n7. File copying and moving:")
	fileCopyMove()

	// 8. File watching and monitoring
	fmt.Println("\n8. File monitoring:")
	fileMonitoring()

	// 9. Working with CSV files
	fmt.Println("\n9. CSV file operations:")
	csvOperations()

	// 10. JSON file operations
	fmt.Println("\n10. JSON file operations:")
	jsonFileOperations()

	// 11. File compression
	fmt.Println("\n11. File compression:")
	fileCompression()

	// 12. File hashing and checksums
	fmt.Println("\n12. File hashing:")
	fileHashing()

	// 13. File locking
	fmt.Println("\n13. File locking:")
	fileLocking()

	// 14. Binary file operations
	fmt.Println("\n14. Binary file operations:")
	binaryFileOperations()

	// 15. File searching and filtering
	fmt.Println("\n15. File searching:")
	fileSearching()

	// 16. Log file operations
	fmt.Println("\n16. Log file operations:")
	logFileOperations()

	// 17. Configuration file handling
	fmt.Println("\n17. Configuration files:")
	configFileHandling()

	// 18. File backup and versioning
	fmt.Println("\n18. File backup:")
	fileBackup()

	// 19. Memory-mapped files
	fmt.Println("\n19. Memory-mapped files:")
	memoryMappedFiles()

	// 20. Advanced file patterns
	fmt.Println("\n20. Advanced file patterns:")
	advancedFilePatterns()

	// Cleanup
	fmt.Println("\n21. Cleanup:")
	cleanup()
}

// 1. Basic file creation and writing
func basicFileOperations() {
	// Create a new file
	file, err := os.Create("example.txt")
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	// Write string to file
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return
	}

	// Write bytes to file
	data := []byte("This is written as bytes\n")
	_, err = file.Write(data)
	if err != nil {
		log.Printf("Error writing bytes: %v", err)
		return
	}

	fmt.Println("File created and written successfully")

	// Append to file
	file2, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file for append: %v", err)
		return
	}
	defer file2.Close()

	_, err = file2.WriteString("Appended line\n")
	if err != nil {
		log.Printf("Error appending to file: %v", err)
		return
	}

	fmt.Println("Text appended to file")
}

// 2. Reading files different ways
func readFileExamples() {
	// Method 1: Read entire file into memory
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}
	fmt.Printf("File content:\n%s", content)

	// Method 2: Read file line by line
	file, err := os.Open("example.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading line by line:")
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file: %v", err)
	}

	// Method 3: Read with buffer
	file2, err := os.Open("example.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file2.Close()

	fmt.Println("Reading with buffer:")
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading: %v", err)
			break
		}
		fmt.Printf("Read %d bytes: %s", n, buffer[:n])
	}
}

// 3. File info and stats
func fileInfoExamples() {
	info, err := os.Stat("example.txt")
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return
	}

	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("File mode: %s\n", info.Mode())
	fmt.Printf("Modification time: %s\n", info.ModTime())
	fmt.Printf("Is directory: %t\n", info.IsDir())

	// Check if file exists
	if _, err := os.Stat("nonexistent.txt"); os.IsNotExist(err) {
		fmt.Println("nonexistent.txt does not exist")
	}

	// File type checks
	if info.Mode().IsRegular() {
		fmt.Println("example.txt is a regular file")
	}
}

// 4. Directory operations
func directoryOperations() {
	// Create directory
	err := os.Mkdir("testdir", 0755)
	if err != nil && !os.IsExist(err) {
		log.Printf("Error creating directory: %v", err)
		return
	}
	fmt.Println("Directory created")

	// Create nested directories
	err = os.MkdirAll("testdir/subdir/nested", 0755)
	if err != nil {
		log.Printf("Error creating nested directories: %v", err)
		return
	}
	fmt.Println("Nested directories created")

	// List directory contents
	entries, err := os.ReadDir("testdir")
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("  %s (size: %d, dir: %t)\n", entry.Name(), info.Size(), entry.IsDir())
	}

	// Walk directory tree
	fmt.Println("Walking directory tree:")
	err = filepath.Walk("testdir", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("  %s\n", path)
		return nil
	})
	if err != nil {
		log.Printf("Error walking directory: %v", err)
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
		return
	}
	fmt.Printf("Current working directory: %s\n", cwd)
}

// 5. File permissions and modes
func filePermissions() {
	// Create file with specific permissions
	file, err := os.OpenFile("restricted.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error creating restricted file: %v", err)
		return
	}
	defer file.Close()

	file.WriteString("This is a restricted file\n")

	// Change file permissions
	err = os.Chmod("restricted.txt", 0644)
	if err != nil {
		log.Printf("Error changing permissions: %v", err)
		return
	}

	// Check permissions
	info, err := os.Stat("restricted.txt")
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return
	}

	fmt.Printf("File permissions: %s\n", info.Mode().Perm())

	// Check if file is readable/writable
	file2, err := os.OpenFile("restricted.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("File is not readable")
	} else {
		fmt.Println("File is readable")
		file2.Close()
	}
}

// 6. Temporary files and directories
func tempFileOperations() {
	// Create temporary file
	tempFile, err := ioutil.TempFile("", "example_*.txt")
	if err != nil {
		log.Printf("Error creating temp file: %v", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up
	defer tempFile.Close()

	fmt.Printf("Temporary file created: %s\n", tempFile.Name())

	// Write to temp file
	_, err = tempFile.WriteString("This is temporary data\n")
	if err != nil {
		log.Printf("Error writing to temp file: %v", err)
		return
	}

	// Create temporary directory
	tempDir, err := ioutil.TempDir("", "example_dir_")
	if err != nil {
		log.Printf("Error creating temp directory: %v", err)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up

	fmt.Printf("Temporary directory created: %s\n", tempDir)

	// Create file in temp directory
	tempFilePath := filepath.Join(tempDir, "temp_file.txt")
	err = ioutil.WriteFile(tempFilePath, []byte("Temp file content"), 0644)
	if err != nil {
		log.Printf("Error creating file in temp dir: %v", err)
		return
	}

	fmt.Printf("File created in temp directory: %s\n", tempFilePath)
}

// 7. File copying and moving
func fileCopyMove() {
	// Copy file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		log.Printf("Error opening source file: %v", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example_copy.txt")
	if err != nil {
		log.Printf("Error creating destination file: %v", err)
		return
	}
	defer destFile.Close()

	// Copy content
	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		return
	}

	fmt.Printf("Copied %d bytes\n", bytesWritten)

	// Move/rename file
	err = os.Rename("example_copy.txt", "example_moved.txt")
	if err != nil {
		log.Printf("Error moving file: %v", err)
		return
	}

	fmt.Println("File moved/renamed successfully")
}

// 8. File watching and monitoring
func fileMonitoring() {
	// Create a file to monitor
	file, err := os.Create("monitored.txt")
	if err != nil {
		log.Printf("Error creating monitored file: %v", err)
		return
	}
	file.Close()

	// Get initial file info
	initialInfo, err := os.Stat("monitored.txt")
	if err != nil {
		log.Printf("Error getting initial file info: %v", err)
		return
	}

	fmt.Printf("Initial file modification time: %s\n", initialInfo.ModTime())

	// Modify the file
	time.Sleep(1 * time.Second)
	file, err = os.OpenFile("monitored.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file for modification: %v", err)
		return
	}
	file.WriteString("Modified content\n")
	file.Close()

	// Check if file was modified
	newInfo, err := os.Stat("monitored.txt")
	if err != nil {
		log.Printf("Error getting new file info: %v", err)
		return
	}

	fmt.Printf("New file modification time: %s\n", newInfo.ModTime())

	if newInfo.ModTime().After(initialInfo.ModTime()) {
		fmt.Println("File was modified!")
	}
}

// 9. Working with CSV files
func csvOperations() {
	// Write CSV file
	csvFile, err := os.Create("data.csv")
	if err != nil {
		log.Printf("Error creating CSV file: %v", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "Age", "City"}
	writer.Write(header)

	// Write data
	records := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "25", "San Francisco"},
		{"Charlie", "35", "Chicago"},
	}

	for _, record := range records {
		writer.Write(record)
	}

	fmt.Println("CSV file written successfully")

	// Read CSV file
	csvFile2, err := os.Open("data.csv")
	if err != nil {
		log.Printf("Error opening CSV file: %v", err)
		return
	}
	defer csvFile2.Close()

	reader := csv.NewReader(csvFile2)
	records2, err := reader.ReadAll()
	if err != nil {
		log.Printf("Error reading CSV file: %v", err)
		return
	}

	fmt.Println("CSV file contents:")
	for i, record := range records2 {
		fmt.Printf("Row %d: %v\n", i, record)
	}
}

// 10. JSON file operations
func jsonFileOperations() {
	// Define struct for JSON data
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}

	// Create sample data
	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "San Francisco"},
		{"Charlie", 35, "Chicago"},
	}

	// Write JSON to file
	jsonFile, err := os.Create("people.json")
	if err != nil {
		log.Printf("Error creating JSON file: %v", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ") // Pretty print
	err = encoder.Encode(people)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		return
	}

	fmt.Println("JSON file written successfully")

	// Read JSON from file
	jsonFile2, err := os.Open("people.json")
	if err != nil {
		log.Printf("Error opening JSON file: %v", err)
		return
	}
	defer jsonFile2.Close()

	var peopleRead []Person
	decoder := json.NewDecoder(jsonFile2)
	err = decoder.Decode(&peopleRead)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	fmt.Println("JSON file contents:")
	for i, person := range peopleRead {
		fmt.Printf("Person %d: %+v\n", i+1, person)
	}
}

// 11. File compression
func fileCompression() {
	// Create a file to compress
	originalFile, err := os.Create("compress_me.txt")
	if err != nil {
		log.Printf("Error creating file to compress: %v", err)
		return
	}

	// Write some repetitive content that compresses well
	content := strings.Repeat("This is a line that will be repeated many times.\n", 100)
	originalFile.WriteString(content)
	originalFile.Close()

	// Get original file size
	originalInfo, _ := os.Stat("compress_me.txt")
	fmt.Printf("Original file size: %d bytes\n", originalInfo.Size())

	// Compress the file
	originalFile, err = os.Open("compress_me.txt")
	if err != nil {
		log.Printf("Error opening file for compression: %v", err)
		return
	}
	defer originalFile.Close()

	compressedFile, err := os.Create("compressed.gz")
	if err != nil {
		log.Printf("Error creating compressed file: %v", err)
		return
	}
	defer compressedFile.Close()

	gzipWriter := gzip.NewWriter(compressedFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, originalFile)
	if err != nil {
		log.Printf("Error compressing file: %v", err)
		return
	}

	// Get compressed file size
	gzipWriter.Close()
	compressedInfo, _ := os.Stat("compressed.gz")
	fmt.Printf("Compressed file size: %d bytes\n", compressedInfo.Size())
	fmt.Printf("Compression ratio: %.2f%%\n", float64(compressedInfo.Size())/float64(originalInfo.Size())*100)

	// Decompress the file
	compressedFile2, err := os.Open("compressed.gz")
	if err != nil {
		log.Printf("Error opening compressed file: %v", err)
		return
	}
	defer compressedFile2.Close()

	gzipReader, err := gzip.NewReader(compressedFile2)
	if err != nil {
		log.Printf("Error creating gzip reader: %v", err)
		return
	}
	defer gzipReader.Close()

	decompressedFile, err := os.Create("decompressed.txt")
	if err != nil {
		log.Printf("Error creating decompressed file: %v", err)
		return
	}
	defer decompressedFile.Close()

	_, err = io.Copy(decompressedFile, gzipReader)
	if err != nil {
		log.Printf("Error decompressing file: %v", err)
		return
	}

	fmt.Println("File compressed and decompressed successfully")
}

// 12. File hashing and checksums
func fileHashing() {
	// Create a file to hash
	file, err := os.Create("hash_me.txt")
	if err != nil {
		log.Printf("Error creating file to hash: %v", err)
		return
	}
	file.WriteString("This content will be hashed")
	file.Close()

	// Calculate MD5 hash
	file, err = os.Open("hash_me.txt")
	if err != nil {
		log.Printf("Error opening file for hashing: %v", err)
		return
	}
	defer file.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Printf("Error calculating hash: %v", err)
		return
	}

	hash := hasher.Sum(nil)
	fmt.Printf("MD5 hash: %x\n", hash)

	// Verify file integrity by comparing hashes
	file2, err := os.Open("hash_me.txt")
	if err != nil {
		log.Printf("Error opening file for verification: %v", err)
		return
	}
	defer file2.Close()

	hasher2 := md5.New()
	io.Copy(hasher2, file2)
	hash2 := hasher2.Sum(nil)

	if string(hash) == string(hash2) {
		fmt.Println("File integrity verified - hashes match")
	} else {
		fmt.Println("File integrity check failed - hashes don't match")
	}
}

// 13. File locking
func fileLocking() {
	// Create a file for locking demonstration
	file, err := os.Create("locked_file.txt")
	if err != nil {
		log.Printf("Error creating file for locking: %v", err)
		return
	}
	defer file.Close()

	file.WriteString("This file will be locked\n")

	// In Go, file locking is platform-specific
	// This is a simplified example - in practice you'd use syscalls
	fmt.Println("File locking is platform-specific")
	fmt.Println("On Unix systems, you can use flock()")
	fmt.Println("On Windows, you can use LockFileEx()")
	fmt.Println("Consider using libraries like github.com/gofrs/flock for cross-platform locking")
}

// 14. Binary file operations
func binaryFileOperations() {
	// Write binary data
	binaryFile, err := os.Create("binary_data.bin")
	if err != nil {
		log.Printf("Error creating binary file: %v", err)
		return
	}
	defer binaryFile.Close()

	// Write different data types
	data := []interface{}{
		uint32(12345),
		float64(3.14159),
		[]byte("binary string"),
	}

	for _, item := range data {
		switch v := item.(type) {
		case uint32:
			bytes := make([]byte, 4)
			bytes[0] = byte(v)
			bytes[1] = byte(v >> 8)
			bytes[2] = byte(v >> 16)
			bytes[3] = byte(v >> 24)
			binaryFile.Write(bytes)
		case float64:
			// Simplified float64 writing - in practice use encoding/binary
			str := fmt.Sprintf("%.5f", v)
			binaryFile.Write([]byte(str))
		case []byte:
			binaryFile.Write(v)
		}
	}

	fmt.Println("Binary data written to file")

	// Read binary data
	binaryFile2, err := os.Open("binary_data.bin")
	if err != nil {
		log.Printf("Error opening binary file: %v", err)
		return
	}
	defer binaryFile2.Close()

	// Read all binary data
	binaryData, err := ioutil.ReadAll(binaryFile2)
	if err != nil {
		log.Printf("Error reading binary data: %v", err)
		return
	}

	fmt.Printf("Read %d bytes of binary data\n", len(binaryData))
	fmt.Printf("First 20 bytes: %v\n", binaryData[:min(20, len(binaryData))])
}

// 15. File searching and filtering
func fileSearching() {
	// Create some test files
	testFiles := []string{"test1.txt", "test2.log", "test3.txt", "data.csv"}
	for _, filename := range testFiles {
		file, _ := os.Create(filename)
		file.WriteString(fmt.Sprintf("Content of %s", filename))
		file.Close()
	}

	// Find all .txt files
	txtFiles, err := filepath.Glob("*.txt")
	if err != nil {
		log.Printf("Error finding txt files: %v", err)
		return
	}

	fmt.Println("Found .txt files:")
	for _, file := range txtFiles {
		fmt.Printf("  %s\n", file)
	}

	// Find files by size
	fmt.Println("Files larger than 10 bytes:")
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Size() > 10 {
			fmt.Printf("  %s (%d bytes)\n", path, info.Size())
		}
		return nil
	})

	// Find files modified in the last hour
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	fmt.Println("Files modified in the last hour:")
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().After(oneHourAgo) {
			fmt.Printf("  %s (modified: %s)\n", path, info.ModTime().Format("15:04:05"))
		}
		return nil
	})
}

// 16. Log file operations
func logFileOperations() {
	// Create a log file
	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error creating log file: %v", err)
		return
	}
	defer logFile.Close()

	// Create a custom logger
	logger := log.New(logFile, "[APP] ", log.LstdFlags|log.Lshortfile)

	// Write log entries
	logger.Println("Application started")
	logger.Println("Processing data...")
	logger.Printf("Processed %d records", 100)
	logger.Println("Application finished")

	fmt.Println("Log entries written to application.log")

	// Read and display log entries
	logContent, err := ioutil.ReadFile("application.log")
	if err != nil {
		log.Printf("Error reading log file: %v", err)
		return
	}

	fmt.Println("Log file contents:")
	fmt.Printf("%s", logContent)

	// Rotate log file (simplified version)
	if err := os.Rename("application.log", "application.log.old"); err == nil {
		fmt.Println("Log file rotated")
	}
}

// 17. Configuration file handling
func configFileHandling() {
	// Create a simple configuration file
	configContent := `# Application Configuration
server_port=8080
database_host=localhost
database_port=5432
debug_mode=true
max_connections=100
`

	err := ioutil.WriteFile("config.txt", []byte(configContent), 0644)
	if err != nil {
		log.Printf("Error writing config file: %v", err)
		return
	}

	// Read and parse configuration
	file, err := os.Open("config.txt")
	if err != nil {
		log.Printf("Error opening config file: %v", err)
		return
	}
	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	fmt.Println("Configuration loaded:")
	for key, value := range config {
		fmt.Printf("  %s = %s\n", key, value)
	}
}

// 18. File backup and versioning
func fileBackup() {
	// Create original file
	originalContent := "Original content version 1.0"
	err := ioutil.WriteFile("document.txt", []byte(originalContent), 0644)
	if err != nil {
		log.Printf("Error creating original file: %v", err)
		return
	}

	// Create backup before modification
	backupName := fmt.Sprintf("document_backup_%s.txt", time.Now().Format("20060102_150405"))
	sourceFile, err := os.Open("document.txt")
	if err != nil {
		log.Printf("Error opening source file: %v", err)
		return
	}
	defer sourceFile.Close()

	backupFile, err := os.Create(backupName)
	if err != nil {
		log.Printf("Error creating backup file: %v", err)
		return
	}
	defer backupFile.Close()

	_, err = io.Copy(backupFile, sourceFile)
	if err != nil {
		log.Printf("Error copying to backup: %v", err)
		return
	}

	fmt.Printf("Backup created: %s\n", backupName)

	// Modify original file
	modifiedContent := "Modified content version 2.0"
	err = ioutil.WriteFile("document.txt", []byte(modifiedContent), 0644)
	if err != nil {
		log.Printf("Error modifying file: %v", err)
		return
	}

	fmt.Println("Original file modified")

	// List all backup files
	backupFiles, err := filepath.Glob("document_backup_*.txt")
	if err != nil {
		log.Printf("Error finding backup files: %v", err)
		return
	}

	fmt.Println("Available backups:")
	for _, backup := range backupFiles {
		info, _ := os.Stat(backup)
		fmt.Printf("  %s (created: %s)\n", backup, info.ModTime().Format("2006-01-02 15:04:05"))
	}
}

// 19. Memory-mapped files
func memoryMappedFiles() {
	// Create a file for memory mapping
	content := make([]byte, 1024)
	for i := range content {
		content[i] = byte(i % 256)
	}

	err := ioutil.WriteFile("mmap_test.dat", content, 0644)
	if err != nil {
		log.Printf("Error creating file for memory mapping: %v", err)
		return
	}

	// Note: Memory mapping in Go requires platform-specific code or external libraries
	// This is a conceptual example
	fmt.Println("Memory-mapped files require platform-specific implementation")
	fmt.Println("Consider using libraries like:")
	fmt.Println("  - github.com/edsrzf/mmap-go")
	fmt.Println("  - golang.org/x/exp/mmap")

	// Read file normally for comparison
	file, err := os.Open("mmap_test.dat")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 10)
	n, err := file.Read(buffer)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}

	fmt.Printf("Read %d bytes: %v\n", n, buffer[:n])
}

// 20. Advanced file patterns
func advancedFilePatterns() {
	// File watching with polling
	fmt.Println("File watching with polling:")
	watchFile := "watch_me.txt"
	ioutil.WriteFile(watchFile, []byte("initial content"), 0644)

	var lastModTime time.Time
	if info, err := os.Stat(watchFile); err == nil {
		lastModTime = info.ModTime()
	}

	// Simulate file watching for a few iterations
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)

		// Modify file
		content := fmt.Sprintf("modified content %d", i+1)
		ioutil.WriteFile(watchFile, []byte(content), 0644)

		// Check for modification
		if info, err := os.Stat(watchFile); err == nil {
			if info.ModTime().After(lastModTime) {
				fmt.Printf("File %s was modified at %s\n", watchFile, info.ModTime().Format("15:04:05"))
				lastModTime = info.ModTime()
			}
		}
	}

	// Atomic file operations
	fmt.Println("\nAtomic file operations:")
	atomicWrite("atomic_test.txt", "This write is atomic")

	// File size monitoring
	fmt.Println("\nFile size monitoring:")
	monitorFile := "growing_file.txt"
	file, _ := os.Create(monitorFile)

	for i := 0; i < 5; i++ {
		file.WriteString(fmt.Sprintf("Line %d\n", i+1))
		file.Sync() // Ensure data is written to disk

		if info, err := os.Stat(monitorFile); err == nil {
			fmt.Printf("File size: %d bytes\n", info.Size())
		}
		time.Sleep(100 * time.Millisecond)
	}
	file.Close()
}

// Helper function for atomic write
func atomicWrite(filename, content string) error {
	tempFile, err := ioutil.TempFile(filepath.Dir(filename), ".tmp_"+filepath.Base(filename))
	if err != nil {
		return err
	}

	_, err = tempFile.WriteString(content)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return err
	}

	err = tempFile.Close()
	if err != nil {
		os.Remove(tempFile.Name())
		return err
	}

	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		os.Remove(tempFile.Name())
		return err
	}

	fmt.Printf("Atomic write completed for %s\n", filename)
	return nil
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 21. Cleanup function
func cleanup() {
	// List all files to clean up
	filesToClean := []string{
		"example.txt", "example_moved.txt", "restricted.txt", "monitored.txt",
		"data.csv", "people.json", "compress_me.txt", "compressed.gz", "decompressed.txt",
		"hash_me.txt", "locked_file.txt", "binary_data.bin", "test1.txt", "test2.log",
		"test3.txt", "application.log", "application.log.old", "config.txt", "document.txt",
		"mmap_test.dat", "watch_me.txt", "atomic_test.txt", "growing_file.txt",
	}

	// Find backup files
	backupFiles, _ := filepath.Glob("document_backup_*.txt")
	filesToClean = append(filesToClean, backupFiles...)

	// Clean up files
	for _, file := range filesToClean {
		if err := os.Remove(file); err == nil {
			fmt.Printf("Cleaned up: %s\n", file)
		}
	}

	// Clean up directories
	dirsToClean := []string{"testdir"}
	for _, dir := range dirsToClean {
		if err := os.RemoveAll(dir); err == nil {
			fmt.Printf("Cleaned up directory: %s\n", dir)
		}
	}

	fmt.Println("Cleanup completed")
}
