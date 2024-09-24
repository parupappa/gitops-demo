package main

import (
	"fmt"
	"log"
	"net/http"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Fibonacci Sequence</title>
	</head>
	<body>
		<h1>Fibonacci Sequence</h1>
		<button id="fibonacciBtn" style="background-color: green; color: white;">数列を表示</button> <div id="result"></div>
		<div id="result"></div>
		<script>
			const fibonacciBtn = document.getElementById('fibonacciBtn');
			const resultDiv = document.getElementById('result');

			fibonacciBtn.addEventListener('click', () => {
				let output = '<ul>';
				for (let i = 0; i < 10; i++) {
					output += '<li>' + fibonacci(i) + '</li>'; // 修正箇所：テンプレートリテラルを通常の文字列連結に変更
				}
				output += '</ul>';
				resultDiv.innerHTML = output;
			});

			function fibonacci(n) {
				if (n <= 1) {
					return n;
				}
				return fibonacci(n - 1) + fibonacci(n - 2);
			}
		</script>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
