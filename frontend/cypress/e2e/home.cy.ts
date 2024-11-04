// https://on.cypress.io/api

describe('Home page test', () => {
  it('visits the app root url', () => {
    cy.viewport(1500,1000)
    cy.visit('/')
    // home page
    cy.get('[cy="welcome-word"]').contains("Your Links, Your Story")
    cy.get('[cy="logout-btn"]').should('not.visible')

    // about page
    cy.get('[cy="about-me"').click()
    cy.contains("2024").should("be.visible")
    cy.get('[cy="me-img"]').should("be.exist")
    cy.get('[cy="login-btn"]').click()

    // login page    
    cy.contains('Register').click()

    // register page
    cy.contains('First Name').type('Doha')
    cy.contains('Last Name').type('Elsawy')
    cy.contains('Email').type('finally@gmail.com')
    cy.contains('Username').type('finally')
    cy.contains('Password').type('1234')
    cy.get('[cy="register-btn"]').click()
    cy.on('window:alert', (alertText) => {
      expect(alertText).to.exist
    })
    cy.get('[cy="password"]').clear()
    cy.contains('Password').type('12345678')
    cy.get('[cy="register-btn"]').click()

    // login page
    cy.get('[cy="login-email"]').type('finally@gmail.com')
    cy.get('[cy="login-password"]').type('12345678')
    cy.get('[cy="login-btn"]').click()
    cy.wait

    // home page
    cy.get('[cy="logout-btn"]').should('exist')
    cy.getCookie("Authorization").should("exist")
    cy.get('[cy="get-start-btn"]').click()

    // link page
    cy.get('[cy="add-link-btn"]').click()
    cy.get('[cy="add-name"]').type("name")
    cy.get('[cy="add-url"]').type("url")
    cy.get('[cy="add-btn"]').click()
    cy.contains("name").should('be.visible')
    cy.get('[cy="live-demo-btn"]').click()
    cy.go(-1)
    cy.get('[cy="link-edit"]').first().click().pause
    cy.get('[cy="edit-name"]').type("edited name")
    cy.get('[cy="edit-url"]').type("edited url")
    cy.get('[ cy="edit-btn"]').click()
    cy.get('[cy="link-delete"]').first().click()
    cy.get('[cy="profile"]').click()

    // profile: edit page
    cy.get('[cy="img"]').selectFile("cypress/e2e/testdata/test-img.jpg")
    cy.get('[cy="bio"]').type("hello, it's me")
    cy.get('[cy="profile-update-btn"]').click()

    // profile: show analysis 
    cy.get('[cy="profile-analysis"]').click()
    cy.contains("Click Count").should('be.visible')
    cy.get('[cy="profile-edit"]').click()
    
    // profile: edit page
    cy.get('[cy="profile-delete-btn"]').click()
  })
})

