---
layout: post
title: Advices
nav_order: 2
has_children: true
has_toc: false
---

{: .text-green-300}
# Advice
{: .fs-9 }

{: .text-green-200}
The piece that you were missing...
{: .fs-6 .fw-300 }

---
 
{: .text-green-300}
## Types of advices
 

{: .text-yellow-300}
### Before 

{: .text-green-200}
#### Description

Advice that must be executed before a function/method invocation.

{: .text-green-200}
#### Use cases

- Trace functions and methods calls.
- Populate arguments with external data before being processed. 
- Normalize arguments before being passed to functions and methods. 
- Arguments validation.
- Security checks.
- ...

[Go to section "Before advices"](/advices/before/)
{: .fs-4 }

---

{: .text-yellow-300}
### Returning

{: .text-green-200}
#### Description
Advice that must be executed after a function/method invocation.

{: .text-green-200}
#### Use cases

- Update fields depending on output response.
- Normalize results. 
- Check errors and manipulate them.
- Throw business errors depending on the results.
- ...

[Go to section "Returning advices"](/advices/returning/)
{: .fs-4 }
---

{: .text-yellow-300}
### Around

{: .text-green-200}
#### Description
Advice that surrounds a function/method invocation.

{: .text-green-200}
#### Use cases

- Implement custom and smart memorize advices.
- Timing your functions and methods. 
- Metric your function and methods.
- ...

[Go to section "Around advices"](/advices/around)
{: .fs-4 }

---

