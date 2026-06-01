# you aren't gonna need it": don't implement something until it is necessary.

## advises against adding functionality until it is necessary

Even if you're sure that you'll need a feature later on, don't implement it now. Usually, it'll turn out either a) you don't need it after all, or b) what you need is quite different from what you foresaw needing earlier.

This doesn't mean you should avoid building flexibility into your code. It means you shouldn't overengineer something based on what you think you might need later on. 

You're working on some class. You have just added some functionality that you need. You realize that you are going to need some other bit of functionality. If you don't need it now, don't add it now. 

### By following YAGNI, you avoid spending time and effort on features that may not be needed right away or at all. adding unnecessary features too early can lead to increased complexity, more bugs, and higher maintenance costs. Instead, you can always iterate and add new features based on actual user feedback and requirements.

Example
Suppose you're working on a basic to-do list application. The core functionality you need to implement includes:
Adding a task.
Marking a task as complete.
Deleting a task.

As you're designing the application, you think about additional features like:
Assigning due dates to tasks.
Setting task priorities.
Creating categories for tasks.

While these features might be useful, they are not part of the initial requirements. According to the YAGNI principle, you should focus on implementing the core functionality first: By following YAGNI, you avoid spending time and effort on features that may not be needed right away or at all. This helps in maintaining a clean and manageable codebase and ensures that the development process remains focused on delivering the essential functionality that users need.