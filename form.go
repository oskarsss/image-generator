package main

import "net/http"

func formHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
		<body>
			<label for="width">Width:</label>
			<input type="range" id="width" name="width" min="1" max="800" value="400" oninput="updateText('width', 'widthValue')" onmouseup="updateImage()">
			<span id="widthValue">400</span><br><br>

			<label for="height">Height:</label>
			<input type="range" id="height" name="height" min="1" max="800" value="400" oninput="updateText('height', 'heightValue')" onmouseup="updateImage()">
			<span id="heightValue">400</span><br><br>

			<label for="red">Red:</label>
			<input type="range" id="red" name="red" min="0" max="255" value="0" oninput="updateText('red', 'redValue')" onmouseup="updateImage()">
			<span id="redValue">0</span><br><br>

			<label for="green">Green:</label>
			<input type="range" id="green" name="green" min="0" max="255" value="0" oninput="updateText('green', 'greenValue')" onmouseup="updateImage()">
			<span id="greenValue">0</span><br><br>

			<label for="blue">Blue:</label>
			<input type="range" id="blue" name="blue" min="0" max="255" value="255" oninput="updateText('blue', 'blueValue')" onmouseup="updateImage()">
			<span id="blueValue">255</span><br><br>

			<label for="shape">Shape:</label>
			<select id="shape" name="shape" onchange="updateImage()">
				<option value="rectangle">Rectangle</option>
				<option value="circle">Circle</option>
				<option value="triangle">Triangle</option>
				<option value="octagram">Octagram</option>
			</select><br><br>

			<img id="generatedImage" src="/generate?width=400&height=400&red=0&green=0&blue=255&shape=rectangle" alt="Generated Image"><br><br>

			<script>
				function updateText(sliderId, spanId) {
					document.getElementById(spanId).innerText = document.getElementById(sliderId).value;
				}
				function updateImage() {
					var width = document.getElementById('width').value;
					var height = document.getElementById('height').value;
					var red = document.getElementById('red').value;
					var green = document.getElementById('green').value;
					var blue = document.getElementById('blue').value;
					var shape = document.getElementById('shape').value;
					document.getElementById('generatedImage').src = '/generate?width=' + width + '&height=' + height + '&red=' + red + '&green=' + green + '&blue=' + blue + '&shape=' + shape;
				}
			</script>
		</body>
	</html>`
	w.Write([]byte(html))
}
