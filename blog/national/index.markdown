---
layout: blogindex
heading: 'Curiosity Hacked'
category: national
---
{% for post in site.categories.[page.category]%}
<article class="box-excerpt">
<a href="#" class="image image-left">
{% if post.headingimg %}
<img src="{{ post.headingimg }} alt="" />
{% else %}
<img src="images/pic04.jpg" alt="" />
{% endif %}
</a>
<div>
<header>
<span class="date">{{ post.date | date_to_string }} </span>
<h3><a href="{{ post.url }}">{{ post.title | strip_html }}</a></h3>
</header>
<p>
{{ post.excerpt | markupify | strip_html | truncatewords: 50 }}
</p>
</div>
</article>

{% endfor %}
