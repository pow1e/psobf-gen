package obfuscate

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func GenerateObfuscatedScript(word string, level string) string {
	if level == "" {
		log.Fatal("请使用-l参数指定模糊等级")
	}
	if level == "all" {
		return genScript(word, "1") + "\n" + genScript(word, "2") + "\n" + genScript(word, "3") + "\n" + genScript(word, "4") + "\n" + genScript(word, "5")
	}
	return genScript(word, level)
}

func genScript(word string, level string) string {
	switch level {
	case "1":
		arg1 := randomString(rand.Intn(20) + 1)
		concatenated := ""
		for _, char := range word {
			concatenated += obfuscateCharacter(char)
		}
		concatenated = strings.TrimSuffix(concatenated, ",")
		return fmt.Sprintf("$%s = $([char[]](%s) -join ''); Invoke-Expression $%s", arg1, concatenated, arg1)
	case "2":
		arg1 := randomString(rand.Intn(20) + 1)
		encoded := base64.StdEncoding.EncodeToString([]byte(word))
		return fmt.Sprintf("$%s = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('%s')); Invoke-Expression $%s", arg1, encoded, arg1)
	case "3":
		arg1 := randomString(rand.Intn(20) + 1)
		arg2 := randomString(rand.Intn(20) + 1)
		encoded := base64.StdEncoding.EncodeToString([]byte(word))
		return fmt.Sprintf("$%s = [System.Convert]::FromBase64String('%s'); $%s = [System.Text.Encoding]::UTF8.GetString($%s); Invoke-Expression $%s", arg1, encoded, arg1, arg2, arg2)
	case "4":
		arg1 := randomString(rand.Intn(20) + 1)
		arg2 := randomString(rand.Intn(20) + 1)
		arg3 := randomString(rand.Intn(20) + 1)
		arg4 := randomString(rand.Intn(20) + 1)
		arg5 := randomString(rand.Intn(20) + 1)
		arg6 := randomString(rand.Intn(20) + 1)
		encoded := compressAndEncode(word)
		return fmt.Sprintf("$%s = '%s'; $%s = [System.Convert]::FromBase64String($%s); $%s = New-Object IO.MemoryStream(, $%s); $%s = New-Object IO.Compression.GzipStream($%s, [IO.Compression.CompressionMode]::Decompress); $%s = New-Object IO.StreamReader($%s); $%s = $%s.ReadToEnd(); Invoke-Expression $%s", arg1, encoded, arg2, arg1, arg3, arg2, arg4, arg3, arg5, arg4, arg6, arg5, arg6)
	case "5":
		arg1 := randomString(rand.Intn(20) + 1)
		arg2 := randomString(rand.Intn(20) + 1)
		fragments := fragmentScript(word)
		return fmt.Sprintf("$%s = @('%s'); $%s = $%s -join ''; Invoke-Expression $%s", arg1, strings.Join(fragments, "','"), arg2, arg1, arg2)
	default:
		log.Fatalf("不支持该模糊处理级别: %s，请使用-l参数指定模糊等级", level)
		return ""
	}
}

func findVariables(script string) []string {
	variables := make(map[string]bool)
	lines := strings.Split(script, "\n")
	for _, line := range lines {
		if strings.Contains(line, "$") {
			parts := strings.Fields(line)
			for _, part := range parts {
				if strings.HasPrefix(part, "$") {
					variables[part] = true
				}
			}
		}
	}
	keys := make([]string, 0, len(variables))
	for key := range variables {
		keys = append(keys, key)
	}
	return keys
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func fragmentScript(script string) []string {
	rand.Seed(time.Now().UnixNano())
	length := len(script)
	fragmentSize := rand.Intn(10) + 10

	var fragments []string
	for i := 0; i < length; i += fragmentSize {
		end := i + fragmentSize
		if end > length {
			end = length
		}

		fragment := script[i:end]
		fragment = strings.ReplaceAll(fragment, "'", "''")
		fragments = append(fragments, fragment)
	}
	return fragments
}

func compressAndEncode(word string) string {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write([]byte(word))
	if err != nil {
		log.Panicf("Failed to compress: %v", err)
	}
	w.Close()
	return base64.StdEncoding.EncodeToString(b.Bytes())
}
func obfuscateVariables(script string) string {
	variables := findVariables(script)
	for _, variable := range variables {
		obfuscated := randomString(len(variable))
		script = strings.ReplaceAll(script, variable, obfuscated)
	}
	return script
}
func obfuscateCharacter(char rune) string {
	switch char {
	case '\n':
		return "\"`n\","
	case '\'':
		return "\"'\","
	case '`':
		return "\"``\","
	case '$':
		return "\"`$\","
	case '(':
		return "\"`(\","
	case ')':
		return "\"`)\","
	case '|':
		return "\"`|\","
	default:
		return fmt.Sprintf("'%s',", string(char))
	}
}
