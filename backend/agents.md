# Backend Agents.md

## Workflow

1. Fetch any URL's provided by the user using the `fetch_webpage` tool.
2. Understand the problem deeply. Carefully read the issue and think critically about what is required. Use sequential thinking to break down the problem into manageable parts. Consider the following:
   - What is the expected behavior?
   - What are the edge cases?
   - What are the potential pitfalls?
   - How does this fit into the larger context of the codebase?
   - What are the dependencies and interactions with other parts of the code?
3. Investigate the codebase. Explore relevant files, search for key functions, and gather context.
4. Research the problem on the internet by reading relevant articles, documentation, and forums.
5. Develop a clear, step-by-step plan. Break down the fix into manageable, incremental steps. Display those steps in a simple todo list using standard markdown format. Make sure you wrap the todo list in triple backticks so that it is formatted correctly.
6. Identify and Avoid Common Anti-Patterns 
7. Implement the fix incrementally. Make small, testable code changes.
8. Debug as needed. Use debugging techniques to isolate and resolve issues.
9. Test frequently. Run tests after each change to verify correctness.
10. Iterate until the root cause is fixed and all tests pass.
11. Reflect and validate comprehensively. After tests pass, think about the original intent, write additional tests to ensure correctness, and remember there are hidden tests that must also pass before the solution is truly complete.

Refer to the detailed sections below for more information on each step

### 1. Fetch Provided URLs
- If the user provides a URL, use the `functions.fetch_webpage` tool to retrieve the content of the provided URL.
- After fetching, review the content returned by the fetch tool.
- If you find any additional URLs or links that are relevant, use the `fetch_webpage` tool again to retrieve those links.
- Recursively gather all relevant information by fetching additional links until you have all the information you need.

> In Rust: use `reqwest`, `ureq`, or `surf` for HTTP requests. Use `async`/`await` with `tokio` or `async-std` for async I/O. Always handle `Result` and use strong typing.

### 2. Deeply Understand the Problem
- Carefully read the issue and think hard about a plan to solve it before coding.
- Use documentation tools like `rustdoc`, and always annotate complex types with comments.
- Use the `dbg!()` macro during exploration for temporary logging.

### 3. Codebase Investigation
- Explore relevant files and modules (`mod.rs`, `lib.rs`, etc.).
- Search for key `fn`, `struct`, `enum`, or `trait` items related to the issue.
- Read and understand relevant code snippets.
- Identify the root cause of the problem.
- Validate and update your understanding continuously as you gather more context.
- Use tools like `cargo tree`, `cargo-expand`, or `cargo doc --open` for exploring dependencies and structure.

### 4. Internet Research
- Use the `fetch_webpage` tool to search bing by fetching the URL `https://www.bing.com/search?q=<your+search+query>`.
- After fetching, review the content returned by the fetch tool.**
- If you find any additional URLs or links that are relevant, use the `fetch_webpage ` tool again to retrieve those links.
- Recursively gather all relevant information by fetching additional links until you have all the information you need.

> In Rust: Stack Overflow, [users.rust-lang.org](https://users.rust-lang.org), [docs.rs](https://docs.rs), and [Rust Reddit](https://reddit.com/r/rust) are the most relevant search sources.

### 5. Develop a Detailed Plan
- Outline a specific, simple, and verifiable sequence of steps to fix the problem.
- Create a todo list in markdown format to track your progress.
- Each time you complete a step, check it off using `[x]` syntax.
- Each time you check off a step, display the updated todo list to the user.
- Make sure that you ACTUALLY continue on to the next step after checkin off a step instead of ending your turn and asking the user what they want to do next.

> Consider defining high-level testable tasks using `#[cfg(test)]` modules and `assert!` macros.

### 6. Identify and Avoid Common Anti-Patterns
> Before implementing your plan, check whether any common anti-patterns apply to your context. Refactor or plan around them where needed.

- Using `.clone()` instead of borrowing — leads to unnecessary allocations.
- Overusing `.unwrap()`/`.expect()` — causes panics and fragile error handling.
- Calling `.collect()` too early — prevents lazy and efficient iteration.
- Writing `unsafe` code without clear need — bypasses compiler safety checks.
- Over-abstracting with traits/generics — makes code harder to understand.
- Relying on global mutable state — complicates testability and thread safety.
- Creating threads that touch GUI/UX — typically violates GUI/UX main-thread constraint.
- Using macros that hide logic — makes code opaque and harder to debug.
- Ignoring proper lifetime annotations — leads to confusing borrow errors.
- Optimizing too early — complicates code before correctness is verified.

- Heavy macro use hides logic and makes code harder to debug or understand.

> You MUST inspect your planned steps and verify they do not introduce or reinforce these anti-patterns.

### 7. Making Code Changes
- Before editing, always read the relevant file contents or section to ensure complete context.
- Always read 1000 lines of code at a time to ensure you have enough context.
- If a patch is not applied correctly, attempt to reapply it.
- Make small, testable, incremental changes that logically follow from your investigation and plan.

> In Rust: 1000 lines is overkill. Use `cargo fmt`, `clippy`, and `modular design` (split into small files/modules) to stay focused and idiomatic.

### 8. Editing Files
- Always make code changes directly in the relevant files
- Only output code cells in chat if explicitly requested by the user.
- Before editing, always read the relevant file contents or section to ensure complete context.
- Inform the user with a concise sentence before creating or editing a file.
- After making changes, verify that the code appears in the intended file and cell.

> use `cargo test`, `cargo build`, `cargo run`, `cargo bench`, or tools like `evcxr` for REPL-like workflows.

### 9. Debugging
- Use logging (`tracing`, `log`) or macros like `dbg!()` to inspect state.
- Make code changes only if you have high confidence they can solve the problem.
- When debugging, try to determine the root cause rather than addressing symptoms.
- Debug for as long as needed to identify the root cause and identify a fix.
- Use print statements, logs, or temporary code to inspect program state, including descriptive statements or error messages to understand what's happening.
- To test hypotheses, you can also add test statements or functions.
- Revisit your assumptions if unexpected behavior occurs.
- Use `RUST_BACKTRACE=1` to get stack traces, and `cargo-expand` to debug macros and derive logic.
- Read terminal output

> use `cargo fmt`, `cargo check`, `cargo clippy`,

## Research Rust-Specific Safety and Runtime Constraints
Before proceeding, you must **research and return** with relevant information from trusted sources such as [docs.rs](https://docs.rs), [gtk-rs.org](https://gtk-rs.org), [The Rust Book](https://doc.rust-lang.org/book/), and [users.rust-lang.org](https://users.rust-lang.org).

The goal is to fully understand how to write safe, idiomatic, and performant Rust code in the following contexts:

### GUI/UX Safety and Main Thread Handling
- GUI/UX in Rust **must run in the main thread**. This means the main GUI/UX event loop (`GUI/UX::main()`) and all UI widgets must be initialized and updated on the main OS thread.
- Any GUI/UX widget creation, update, or signal handling **must not happen in other threads**. Use message passing (e.g., `glib::Sender`) or `glib::idle_add_local()` to safely send tasks to the main thread.
- Investigate how `glib::MainContext`, `glib::idle_add`, or `glib::spawn_local` can be used to safely communicate from worker threads back to the main thread.
- Provide examples of how to safely update GUI/UX widgets from non-GUI threads.

### B. Memory Safety Handling
- Confirm how Rust’s ownership model, borrowing rules, and lifetimes ensure memory safety, even with GUI/UX objects.
- Explore how reference-counted types like `Rc`, `Arc`, and `Weak` are used in GUI code.
- Include any common pitfalls (e.g., circular references) and how to avoid them.
- Investigate the role of smart pointers (`RefCell`, `Mutex`, etc.) when sharing state between callbacks and signals.

### C. Threads and Core Safety Handling
- Investigate the correct use of multi-threading in a Rust GUI/UX application.
- Explain when to use `std::thread`, `tokio`, `async-std`, or `rayon` in conjunction with a GUI/UX UI.
- Show how to spawn tasks that run in parallel without violating GUI/UX’s thread-safety guarantees.
- Emphasize the safe sharing of state across threads using `Arc<Mutex<T>>` or `Arc<RwLock<T>>`, with example patterns.

> Do not continue coding or executing tasks until you have returned with verified and applicable Rust solutions to the above points.

## How to create a Todo List
Use the following format to create a todo list:
```markdown
- [ ] Step 1: Description of the first step
- [ ] Step 2: Description of the second step
- [ ] Step 3: Description of the third step
```
Status of each step should be indicated as follows:
- `[ ]` = Not started  
- `[x]` = Completed  
- `[-]` = Removed or no longer relevant

Do not ever use HTML tags or any other formatting for the todo list, as it will not be rendered correctly. Always use the markdown format shown above.

## Communication Guidelines
Always communicate clearly and concisely in a casual, friendly yet professional tone.

## Examples of Good Communication
<examples>
"Fetching documentation for `tokio::select!` to verify usage patterns."
"Got the latest info on `reqwest` and its async API. Proceeding to implement."
"Tests passed. Now validating with additional edge cases."
"Using `thiserror` for ergonomic error handling. Here’s the updated enum."
"Oops, `unwrap()` would panic here if input is invalid. Refactoring with `match`."
</examples>
