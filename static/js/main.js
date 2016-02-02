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
    switch ((stepClass - 0)) {
      case 1:
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      case 2:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
        break;
      case 3:
        $("#title").text("模块化");
        $("#content1").text("内置强大的模块，包括 Session、缓存操作、日志记录、配置解析、性能监控、上下文操作、ORM 模块、请求模拟等强大的模块，足以支撑你任何的应用。");
        break;
      case 4:
        $("#title").text("高性能");
        $("#content1").text("采用Go 原生的 http 包来处理请求，goroutine 的并发效率足以应付大流量的 Web 应用和 API 应用，目前已经应用于大量高并发的产品中。");
        break;
      case 5:
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      default:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
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
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      case 2:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
        break;
      case 3:
        $("#title").text("模块化");
        $("#content1").text("内置强大的模块，包括 Session、缓存操作、日志记录、配置解析、性能监控、上下文操作、ORM 模块、请求模拟等强大的模块，足以支撑你任何的应用。");
        break;
      case 4:
        $("#title").text("高性能");
        $("#content1").text("采用Go 原生的 http 包来处理请求，goroutine 的并发效率足以应付大流量的 Web 应用和 API 应用，目前已经应用于大量高并发的产品中。");
        break;
      case 5:
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      default:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
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
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      case 2:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
        break;
      case 3:
        $("#title").text("模块化");
        $("#content1").text("内置强大的模块，包括 Session、缓存操作、日志记录、配置解析、性能监控、上下文操作、ORM 模块、请求模拟等强大的模块，足以支撑你任何的应用。");
        break;
      case 4:
        $("#title").text("高性能");
        $("#content1").text("采用Go 原生的 http 包来处理请求，goroutine 的并发效率足以应付大流量的 Web 应用和 API 应用，目前已经应用于大量高并发的产品中。");
        break;
      case 5:
        $("#title").text("简单化");
        $("#content1").text("RESTful 支持、MVC 模型，可以使用 bee 工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。");
        break;
      default:
        $("#title").text("智能化");
        $("#content1").text("支持智能路由、智能监控，可以监控 QPS、内存消耗、CPU 使用，以及 goroutine 的运行状况，让您的线上应用尽在掌握。");
        break;
    }


    $("#bbp").fadeIn(800);


    currentNum += 1;
    if (currentNum == 7) {
      currentNum = 1;
    }


  }

});
