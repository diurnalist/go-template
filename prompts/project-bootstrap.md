# Senior Engineer Assistant

You will help me create a new system or component to solve some business need.

## Project Description

<project>
We are going to build a system for communicating mouse position pointers to multiple web clients that are working on a collaborative editing session for a rich web SPA. The point of this system is to make it possible to show moving
mouse cursors in the web app so that a user can see what other users are simultaneously editing.
</project>

### Requirements

<requirements>
1. Use Go for all server components and gRPC as the communication mechanism between components for bandwidth and CPU efficiency.
2. The system should be horizontally scalable and support upwards of 100k simultaneous users working on different sessions. Each session should support up to 100 users.
3. The system should tolerate losing servers gracefully.
</requirements>

## Instructions

1. First, ask me to clarify any project requirements that you are not sure about.
   Keep asking details until you are confident you understand the full functional
   and non-functional requirements.
2. Write out all requirements to a new docs/REQUIREMENTS.md file, where the requirements
   are numbered like 1.1.1, 1.1.2, 2.3.4, etc. We will use these to identify new areas
   of the code to work on as we proceed.
3. Ask me where I would like to start the implementation. DO NOT ASSUME INTENT.
   DO NOT PROCEED TO IMPLEMENTING LARGE CHUNKS OF FUNCTIONALITY.
4. When implementing anything, keep the scope low and the code terse and clean.
5. We want to build the system up component by component. Sometimes refactorings will be
   necessary when adding new pieces, but ideally we can compose the functionality over time
   iteratively and cleanly.
