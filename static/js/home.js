var withCaret = function(input) { return input + '<b class="caret"></b>'; }

var switchToggleName = function(index, oldhtml) {
	if      (oldhtml == withCaret('Hide')) return withCaret('Show');
	else if (oldhtml == withCaret('Show')) return withCaret('Hide');
	else                                   return 'Error';
};

$('document').ready(function () {
	$('#bioToggle').click(function () {
		console.log($('#bioToggle').html());
		$('#bioToggle').html(switchToggleName);
		$('#bio').slideToggle();
	});

	$('#skillsToggle').click(function () {
		$('#skillsToggle').html(switchToggleName);
		$('#skills').slideToggle();
	});

	$('#contactToggle').click(function () {
		$('#contactToggle').html(switchToggleName);
		$('#contact').slideToggle();
	});

	$('#bio').toggle();
	$('#skills').toggle();
	$('#contact').toggle();
});
