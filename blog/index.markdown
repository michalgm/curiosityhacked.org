---
layout: default
---
###Curiosity Hacked 
<ul class="more">
  {% for post in site.tags.front limit: 4 %}
    <li>
      {{ post.date | date_to_string }} <a href="{{ post.url }}">{{ post.title }}</a>
      {{ post.excerpt | markupify | strip_html }}
    </li>
  {% endfor %}
</ul>

{% capture cats %}
{% for cat in site.categories %}
{{ cat | first }}
{% endfor %}
{% endcapture %}

{% assign cat =  cats |remove: "blog" | split:' ' |  sort  %}

{% for q in cat %}
<li><a href="/blog/{{q}}">{{q}}</a>
{% endfor %}

