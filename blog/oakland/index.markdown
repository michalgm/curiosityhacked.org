---
layout: blogindex
category: oakland
heading: 'Curiosity Hacked Oakland Lab'
---
{% for post in site.categories.[page.category]%}
<article class="box excerpt">
<a href="#" class="image left"><img src="images/pic04.jpg" alt="" /></a>
<div>
<header>
<span class="date">{{ post.date | date_to_string }} </span>
<h3><a href="{{ post.url }}">{{ post.title }}</a></h3>
</header>
<p> {{ post.excerpt | markupify | strip_html }}
</div>
</article>

  {% endfor %}
