---
layout: post
title: Advices
parent: Guides & Tutorials
nav_order: 1
has_children: true
---

{: .text-blue-300}
# Understanding the advices
 
{: .text-blue-200}
## Types of advices

First of all, we'll go through the supported types of advices by Goa. 

{: .text-blue-100}
### Before 

{: .text-blue-000}
#### Description

Advice that executes before a function/method invocation.

{: .text-blue-000}
#### Use cases

- Trace functions and methods calls.
- Populate a field in an object depending on other values. 
- Normalize arguments before being passed to functions and methods. 
- Arguments validation.
- Security checks.
- ...

[Go to section "Before advices"](/guides/advices/before/){: .text-blue-300 }

{: .text-blue-100}
### Returning

{: .text-blue-000}
#### Description
Advice that executes after a function/method invocation.

{: .text-blue-000}
#### Use cases

- Update field depending on output response.
- Normalize results. 
- Check errors and manipulate them.
- Throw business errors depending on the result.
- ...

[Go to section "Returning advices"](/guides/advices/returning/){: .text-blue-300 }

{: .text-blue-100}
### Around

{: .text-blue-000}
#### Description
Advice that surrounds a function/method invocation.

{: .text-blue-000}
#### Use cases

- Implement custom and smart memorize advices.
- Timing your functions and methods. 
- Metric your function and methods.
- ...

[Go to section "Around advices"](/guides/advices/around/){: .text-blue-300 }

---

{: .text-blue-200}
## Goa context

When working with advices is necessary to know the details from the 
functions or methods which are being intercepted. 

That's,  why Goa provides us with **GoaContext**. The GoaContext will be accessible from the different implemented functions
by our advices. It not only provides us with several methods to obtain information from the intercepted functions 
but also with manipulate params and results.


[Go to section "Goa Context"](/guides/advices/context/){: .text-blue-300 }
