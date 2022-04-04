package main

const (
	start string = "<!DOCTYPE html><html lang=\"ru\">"
	head  string = `
					<head>
					<title>GO dev</title>
					<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
					<link rel=stylesheet href=https://cdn.jsdelivr.net/npm/pretty-print-json@1.2/dist/pretty-print-json.css>
					<script src=https://cdn.jsdelivr.net/npm/pretty-print-json@1.2/dist/pretty-print-json.min.js></script>
					</head>`
	form string = `
					<form enctype="application/json" action="/" method="post" _lpchecked="1">
					<label for="search">Поиск</label>
					<input type="text" name="value" id="search" value="">
					<button type="submit">поиск</button></form>
					<pre id="account" class="json-container"></pre>`
	end    string = "</html>"
	script string = `
					<script>$(document).ready(function()
					{
						$( "form" ).submit(function( event ) {
							event.preventDefault();
							let val = $( "input[name='value']" ).val()
							$.ajax({
								url: "/", 
								dataType: "json",
								contentType: "application/json; charset=UTF-8",
								method: "POST",
								data: JSON.stringify({value: val}),
								success: function(data) {
									$(".json-container").html(prettyPrintJson.toHtml(data));
								} 
							});
						});
					});
					</script>				
					`
)

func getHTML() string {
	return start + head + form + script + end
}
