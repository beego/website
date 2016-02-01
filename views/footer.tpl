{{define "footer"}}
<footer id="footer">
  <ul>
    <li>
      <h2>GitHub</h2>
      <a target="_blank " href="https://github.com/astaxie/beego">仓库</a>
      <a target="_blank" href="https://github.com/beego/bee">Bee工具</a>
      <hr>
      <a class="github-button" href="https://github.com/astaxie/beego" data-icon="octicon-star" data-count-href="/astaxie/beego/stargazers" data-count-api="/repos/astaxie/beego#stargazers_count" data-count-aria-label="# Beego on GitHub" aria-label="Star astaxie/beego on GitHub">Star</a>
      <a class="github-button" href="https://github.com/astaxie/beego/fork" data-icon="octicon-repo-forked" data-count-href="/astaxie/beego/network" data-count-api="/repos/astaxie/beego#forks_count" data-count-aria-label="# Beego on GitHub" aria-label="Fork astaxie/beego on GitHub">Fork</a>
    </li>
    <li>
      <h2>我们</h2>
      <a target="_blank" href="/blog">官方博客 - Beego</a>
      <a target="_blank" href="#">反馈和建议</a>
      <a target="_blank" href="#">讨论</a>
      <a target="_blank" href="#">报告 Bug</a>
    </li>
    <li>
      <h2>其他</h2>
      {{if eq .Lang "zh-CN"}}
      <a target="_blank" href="/donate">捐赠我们</a>
      {{end}}
      {{if eq .Lang "en-US"}}
      <a target="_blank" href="/donate"> Donate Us</a>
      {{end}}
      <div class="dropup choose-lang-container">
        <button class="btn dropdown-toggle choose-lang" type="button" id="choose-language" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          Language:English <i class="caret"></i>
        </button>
        <div class="dropdown-menu choose-lang-menu" aria-labelledby="choose-language">
          <a class="dropdown-item choose-lang-item" href="">简体中文</a>
        </div>
      </div>

    </li>
    <li>
      <div>版权所有 © 2012-2014 Beego 授权许可遵循 apache 2.0 licence</div>
      <div>Logo由 Tengfei 设计</div>
      <div>文档版本：1.6.0</div>
    </li>
  </ul>

  <script src='/static/js/jquery-2.1.4.min.js'></script>
  <script src='/static/js/tether.min.js'></script>
  <script src='/static/js/bootstrap.min.js'></script>
  <!-- for github -->
  <script async defer id="github-bjs" src="https://buttons.github.io/buttons.js"></script>
</footer>
{{end}}
