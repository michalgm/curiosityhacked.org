---
layout: default
---
###Curiosity Hacked 
<ul class="more">
  {% for post in site.tags.front %}
    <li style="display:none;">
      {{ post.date | date_to_string }} <a href="{{ post.url }}">{{ post.title }}</a>
      {{ post.excerpt | markupify | strip_html }}
    </li>
  {% endfor %}
</ul>
<script type="text/javascript">
  $("ul.more").each(function() {
    $("li:lt(4)", this).show();
    });
</script>
