var switchToggleName = function(index, oldhtml) {
	if      (oldhtml == 'Hide') return 'Show';
	else if (oldhtml == 'Show') return 'Hide';
	else                        return 'Error';
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

	$('#bio').toggle();
	$('#skills').toggle();
});
