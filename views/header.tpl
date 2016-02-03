{{define "header"}}
<header id="header" class="clearfix">
  <div class="tips-container">
    <a class="tips" href="#"><span>1.6.0</span> Released，<span>1.6.0 Document</span> has been updated</a>
  </div>

  <a class="logo" href="#">
    <img src="/static/img/beego_purple.png"> beego
  </a>
  <div class="search">
    <div id="autoComplete"></div>
  </div>
  <nav class="nav nav-right" style="display: block;">
    <span class="bar"></span>
    <ul>
      {{range .HomeNav}}
      <li>
        <a href="{{.URL}}" class="{{.I18n}}">{{$.I18n.Tr .I18n}}</a>
      </li>
      {{end}}
    </ul>
  </nav>
  <div class="nav-phone-icon"></div>
</header>
{{end}}
