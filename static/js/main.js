(function($){
	$(document).on('click', '.lang-changed', function(){
		var $e = $(this);
		var lang = $e.data('lang');
    var v = $.cookie('JsStorage');
		$.cookie('lang', lang, {path: '/', expires: 365});
		window.location.reload();
	});
})(jQuery);
