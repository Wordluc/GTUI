# Terminal Interface Framework in Go
<img width="2880" height="1590" alt="image" src="https://github.com/user-attachments/assets/3c347ce5-0b4e-4d8d-b7c7-5835f9a7956f" />
This framework allows you to build terminal applications in Go without worrying about the complexities of rendering and event calling. It automatically handles screen refreshing when necessary, optimizing performance.
To see a "demo",you can launch `go run Smoke_Test/main.go`
### Features:

- **Element Support:**  
  The framework supports various UI elements, ranging from simple buttons to textboxes with full copy-and-paste functionality.
  
- **Performance Optimization:**  
  The framework efficiently refreshes the screen when needed to ensure optimal performance during user interactions.

- **Layering System:**  
  A full layering system lets you move UI elements back and forth, providing greater flexibility in UI design.
  
- **Event System:**  
  The event system enables asynchronous interaction with the core of the framework, allowing you to trigger events without waiting for the main loop. This decision was made to optimize performance and reduce blocking during user interactions.
### Showcase

To demonstrate the capabilities of this framework, a **TODO app** has been created: https://github.com/Wordluc/Fedi


