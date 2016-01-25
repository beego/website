{{define "header"}}
<header id="header" class="clearfix">
  <div class="tips-container">
    <a class="tips" href="#"><span>1.6.0</span>已发布，<span>1.6.0文档</span>已更新</a>
  </div>

  <a class="logo" href="#">
    <img src="/static/img/beego_purple.png"> beego
  </a>
  <div class="search">
    <div id="autoComplete"></div>
  </div>
  <nav class="nav" style="display: block;">
    <span class="bar"></span>
    <ul>
      <li class="">
        <a href="#">首页</a>
      </li>
      <li class="">
        <a href="#">快速入门</a>
      </li>
      <li class="current">
        <a href="#">开发文档</a>
      </li>
      <li class="">
        <a href="#">开发者社区</a>
      </li>
      <li>
        <a href="#">视频教程</a>
      </li>
      <li class="">
        <a href="#">产品案例</a>
      </li>
      <li class="">
        <a href="#">官方博客</a>
      </li>
    </ul>
  </nav>
  <div class="nav-phone-icon"></div>
</header>
{{end}}
