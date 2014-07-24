$('document').ready(function () {
	var converter = new Showdown.converter();

	$('#bodyInput').keyup(function () {
		$('#dynamicMarkdown').html(converter.makeHtml($('#bodyInput').val()));
	});
});
