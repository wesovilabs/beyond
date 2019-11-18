---
layout: page
title: About
permalink: /about/
nav_order: 8
---

{: .text-green-300}
# About Goa
{: .fs-9 }

Goa (Golang Oriented to Aspects) is a Golang library inspired in Spring AOP Framework.

The main goals of Goa are:


- It provides Golang developers with a mechanism to **code under AOP paradigm**.
- You can build useful advices and **share with other developers**. at time you can also take advantages of existing 
ones.
- It's completely opensource and there's **not a company behind this project**.
 
---

{: .text-green-300}
## About me

Hey folks! My name is Ivan and I could describe myself like an open source lover who enjoys doing some researches and 
sharing the knowledge with others.

Since I've got experience coding with several programming languages, and working under different paradigms, I 
enjoy doing libraries that provide Go with utilities existing in other languages. 

Any comment, feedback or suggestion will be appreciated. Feel free to drop me an email at `ivan.corrales.solera@gmail.com`
or reach me at any of these social networks. 


{% include share.html %} 

---

{: .text-green-300}
## License
Goa is distributed by an [MIT license](https://github.com/wesovilabs/goa/tree/master/LICENSE.md).

{: .text-green-300}
## Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other 
method with the owners of this repository before making a change. 

Read more about becoming a contributor in our [Goa repository](https://github.com/wesovilabs/goa/tree/master/contributing.md).

{: .text-yellow-300}
### Thank you to the contributors of Goa!

<ul class="list-style-none">
{% for contributor in site.github.contributors %}
  <li class="d-inline-block mr-1">
     <a href="{{ contributor.html_url }}"><img src="{{ contributor.avatar_url }}" width="32" height="32" alt="{{ contributor.login }}"/></a>
  </li>
{% endfor %}
</ul>

{: .text-green-300}
## Code of Conduct
Goa is committed to fostering a welcoming community.

View Goa Code of Conduct on  [Goa repository](https://github.com/wesovilabs/goa/tree/master/CODE_OF_CONDUCT.md).
