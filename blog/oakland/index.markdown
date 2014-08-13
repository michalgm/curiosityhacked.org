---
layout: default
category: oakland
---
###Curiosity Hacked Oakland Lab
<ul>
  {% for post in site.categories.[page.category]%}
    <li>
      {{ post.date | date_to_string }} <a href="{{ post.url }}">{{ post.title }}</a>
      {{ post.excerpt | markupify | strip_html }}
    </li>
  {% endfor %}
</ul>
