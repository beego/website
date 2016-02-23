$(document).ready(function() {
  $("#steps-menu a").click(function(event) {
    // Replaces main content
    event.preventDefault();
    $(this).parent().addClass("is-active");
    $(this).parent().siblings().removeClass("is-active");
    var step = $(this).attr("href");
    $(".step-content").not(step).css("display", "none");
    $(step).fadeToggle();

    // Rotates the wheel
    $("#steps-menu").removeClass();
    var stepClass = step.charAt(6);
    $("#steps-menu").addClass("step-" + stepClass);
    switch ((stepClass-0)) {
      case 1:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      case 2:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
      case 3:
        $("#title").text("MODULAR");
        $("#content1").text("With powerful built-in modules including session control, caching, logging, configuration parsing, performance supervising, context handling, ORM supporting, and requests simulating. You get the powerful foundation for any type of applications.");
        break;
      case 4:
        $("#title").text("HIGH PERFORMANCE");
        $("#content1").text("With native Go http package to handle the requests and the efficient concurrence of goroutine. Your beego applications can handle massive trafic as beego are doing in many productions.");
        break;
      case 5:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      default:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
    }
    currentNum = step.substr(6, 1) - 0 + 1;

  });


  // Read more links
  $("#read-more a").click(function() {
    event.preventDefault();
    var readMore = $(this).attr("href");
    $(readMore).click();
    switch ((readMore.substr(5, 1) - 0)) {
      case 1:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      case 2:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
      case 3:
        $("#title").text("MODULAR");
        $("#content1").text("With powerful built-in modules including session control, caching, logging, configuration parsing, performance supervising, context handling, ORM supporting, and requests simulating. You get the powerful foundation for any type of applications.");
        break;
      case 4:
        $("#title").text("HIGH PERFORMANCE");
        $("#content1").text("With native Go http package to handle the requests and the efficient concurrence of goroutine. Your beego applications can handle massive trafic as beego are doing in many productions.");
        break;
      case 5:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      default:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
    }
    currentNum = readMore.substr(5, 1) - 0 + 1;

  });


  //slideshow style interval
  var autoSwap = setInterval(swap, 5000);
  var currentNum = 2;
  //pause slideshow and reinstantiate on mouseout
  $('#read-more a, #steps-menu a').hover(
    function() {
      clearInterval(autoSwap);
    },
    function() {
      autoSwap = setInterval(swap, 5000);
    });

  //swap images function
  function swap(action) {
    $("#bbp").fadeOut(100);

    $("#steps-menu li a:eq(" + (currentNum - 1) + ")").parent().addClass("is-active");
    $("#steps-menu li a:eq(" + (currentNum - 1) + ")").parent().siblings().removeClass("is-active");
    var step = $("#steps-menu li a:eq(" + (currentNum - 1) + ")").attr("href");
    $(".step-content").not(step).css("display", "none");
    $(step).fadeToggle();

    // Rotates the wheel
    $("#steps-menu").removeClass();
    var stepClass = step.charAt(6);
    $("#steps-menu").addClass("step-" + stepClass);

    switch (currentNum) {
      case 1:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      case 2:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
      case 3:
        $("#title").text("MODULAR");
        $("#content1").text("With powerful built-in modules including session control, caching, logging, configuration parsing, performance supervising, context handling, ORM supporting, and requests simulating. You get the powerful foundation for any type of applications.");
        break;
      case 4:
        $("#title").text("HIGH PERFORMANCE");
        $("#content1").text("With native Go http package to handle the requests and the efficient concurrence of goroutine. Your beego applications can handle massive trafic as beego are doing in many productions.");
        break;
      case 5:
        $("#title").text("EASY TO USE");
        $("#content1").text("With RESTful support, MVC model, and use bee tool to build your apps quickly with features including code hot compile, automated testing, and automated packing and deploying.");
        break;
      default:
        $("#title").text("INTELLIGENT");
        $("#content1").text("With intelligent routing and monitoring, it's able to monitor your QPS, memory and CPU usages, and goroutine status. It provides you the fully control of your online apps.");
        break;
    }


    $("#bbp").fadeIn(800);


    currentNum += 1;
    if (currentNum == 7) {
      currentNum = 1;
    }


  }

});
