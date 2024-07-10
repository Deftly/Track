# Overview

## The Rise of SPAs
- Websites used to be rendered on the server using data and templates. The resulting html, css, and Javascript code was then sent to the client that requested the page and rendered using the browser. 
- The Javascript that was added to these pages was to provide very simple dynamics like small animations and hover effects.
- Eventually though more and more Javascript code was added which led to single-page-applications(SPAs). Basically these are webpages that are rendered on the client, not the server.
- These web applications still need data, which typically comes from the backend through some form of API.
- The applications consume API data and renders a screen for each view of the application, making it feel like a native desktop application. It allows you to click on links, fill out forms and perform other actions without the page ever reloading.

## SPAs with Vanilla Javascript?
- Frontend web applications are all about handling data and displaying that data through a user interface.
- This means the most important task of any web application is to keep the UI in sync with the data as it changes overtime. And this can actually be a very hard problem to solve.
- Complex UIs can have many different data fields, which we refer to as state data. These fields can be inter-connected, and changes to one can conditionally require changes in other state data. This makes it very hard to create complex UIs using vanilla Javascript for the following reasons:
  - Requires lots of direct DOM manipulation and traversing(imperative)
  - Data (state) is often scattered across the DOM and shared across the entire app, rather than being stored in a central location which makes it hard to understand what the current state is and leads to bugs

## Why do Frontend Frameworks Exist?
- Javascript frameworks exist because keeping a UI in sync with data is very difficult. Frameworks abstract away many of these problems and make it easier to build UIs for developers by allowing them to just focus on the data and building the UI itself.
- Different frameworks have different approaches to doing this, but they all aim to keep the UI and state data in sync over time.
- They also often enforce a "correct" way of structuring and writing code.

## What is React?
- React is a declarative, component-based, state-driven JavaScript library for building user interfaces. It's maintained by Facebook and a community of individual developers and companies.

### Based on Components
- Components are the building blocks of UIs in React. What React does is take a collection of components and render them on a webpage
- Complex UIs are built by combining multiple components together. Another advantage of this approach is the components can be reused throughout the application.
- Components can be as small as a button or as large as an entire page.

### Declarative
- React uses a declarative syntax called JSX. This is used to describe what components should look like and how they work, based on the current state.
- React is an abstraction away from the DOM, meaning we never have to interact with it directly as we would with vanilla Javascript.

### State-driven
- React applications are built around the concept of state, which represents the current condition of the application's data.
- Initially, React renders the application based on the initial state, using the components we've created.
- When an event occurs (like a button click or data fetch), it may trigger a state change.
- When state changes occur, we use React's `setState` method (for class components) or state update functions (for functional components with hooks) to update the state.
- React then automatically re-renders the affected components to reflect the latest state, ensuring the UI is always in sync with the underlying data.
- This approach allows for efficient updates, as React only re-renders the parts of the UI that are affected by the state change, not the entire application.
