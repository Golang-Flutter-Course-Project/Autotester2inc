#  User Acceptance Tests — InnoTest

##  MVP-0: UI Mockup and Basic Flow

**Goal:** Validate interface layout and user flow without working logic.

### UAT-001: Interface elements are visible
- **GIVEN** the user opens the prototype  
- **WHEN** the interface is displayed  
- **THEN** the user sees input fields for URL, test conditions, and a "+" button

### UAT-002: Add new condition field
- **GIVEN** the user clicks the "+" button  
- **WHEN** the interface is active  
- **THEN** a new input field appears for entering a test condition

---

##  MVP-1: Core Functionality

**Goal:** Ensure basic interaction between frontend and backend, and DOM parsing logic.

### UAT-101: Submit a valid URL
- **GIVEN** the user enters a valid URL  
- **WHEN** the user clicks “Check”  
- **THEN** the system sends the data to the backend and displays the result

### UAT-102: Add a condition and verify
- **GIVEN** the user inputs a condition like “Is there a 'Login' button?”  
- **WHEN** the test is run  
- **THEN** the system parses the DOM and displays ✅ or ❌

### UAT-103: Dynamically add multiple conditions
- **GIVEN** the user clicks the "+" button  
- **WHEN** the interface is loaded  
- **THEN** a new condition input field is added

---

##  MVP-2: AI Integration and Automation

**Goal:** Validate templates and Selenium integration.

### UAT-201: Load a test template
- **GIVEN** the user selects a template like “Authorization”  
- **WHEN** the template is applied  
- **THEN** the condition fields auto-fill with predefined checks


### UAT-202: Execute tests via Selenium WebDriver
- **GIVEN** all test steps are filled in  
- **WHEN** the user launches the Selenium execution  
- **THEN** the automation runs and test results are shown

---

##  MVP-3: Export and Usability Testing

**Goal:** Validate result export and end-user usability testing.

### UAT-301: Export test results to PDF
- **GIVEN** the user finishes testing  
- **WHEN** the user clicks “Export to PDF”  
- **THEN** a downloadable PDF with all test results is generated

### UAT-302: UI/UX usability session
- **GIVEN** the user participates in usability testing  
- **WHEN** the user completes 3–4 common tasks  
- **THEN** the steps are clear and no external help is needed

