/// https://on.cypress.io/api

describe('Home page test', () => {
  Cypress.on('uncaught:exception', (err, runnable) => {
    // returning false here prevents Cypress from failing the test
    return false
  })
  
  it('visits the app root url', () => {
    cy.viewport(1500,1000)
    cy.visit("/")
    // home page
    cy.get('[cy="welcome-word"]').contains("Your Links, Your Story")
    cy.get('[cy="logout-btn"]').should('not.visible')

    // search function for not exist user
    cy.get('[cy="search-input"]').type("not exist user")
    cy.get('[cy="search-btn"]').click()
    cy.wait
    cy.contains("NOT FOUND :{").should("be.visible")


    // about page
    cy.get('[cy="about-me"').click()
    cy.contains("2024").should("be.visible")
    cy.get('[cy="me-img"]').should("be.exist")
    cy.get('[cy="login-btn"]').click()

    // login page    
    cy.contains('Register').click()

    // register page
    cy.get('[cy="first-name"]').type('Doha')
    cy.get('[cy="last-name"]').type('Elsawy')
    cy.get('[cy="email"]').type('finally@gmail.com')
    cy.get('[cy="username"]').type('finally')
    cy.get('[cy="password"]').type('1234')
    cy.get('[cy="register-btn"]').click()
    cy.on('window:alert', (alertText) => {
      expect(alertText).to.exist
    })
    cy.get('[cy="password"]').clear()
    cy.get('[cy="password"]').type('12345678')
    cy.get('[cy="register-btn"]').click()

    // login page
    cy.get('[cy="login-email"]').type('notexist@gmail.com')
    cy.on('window:alert', (alertText) => {
      expect(alertText).to.exist
    })
    cy.get('[cy="login-email"]').clear()
    cy.get('[cy="login-email"]').type('finally@gmail.com')
    cy.get('[cy="login-password"]').type('notCorrect')
    cy.on('window:alert', (alertText) => {
      expect(alertText).to.exist
    })
    cy.get('[cy="login-password"]').clear()
    cy.get('[cy="login-password"]').type('12345678')
    cy.get('[cy="login-btn"]').click()
    cy.wait

    // home page
    cy.get('[cy="logout-btn"]').should('exist')
    cy.getCookie("Authorization").should("exist")
    cy.get('[cy="get-start-btn"]').click()

    // search function for exist user
    cy.get('[cy="search-input"]').type("finally")
    cy.get('[cy="search-btn"]').click()
    cy.contains("finally").should("be.visible")

    // link page
    cy.get('[cy="link-tree"]').click()
    cy.get('[cy="add-link-btn"]').click()
    cy.get('[cy="add-name"]').type("name")
    cy.get('[cy="add-url"]').type("url")
    cy.get('[cy="add-btn"]').click()
    cy.contains("name").should('be.visible')
    cy.get('[cy="live-demo-btn"]').click()
    cy.contains("name").should('be.visible')
    cy.contains("url").should('be.visible')
    cy.go(-1)
    cy.get('[cy="link-edit"]').first().click().pause
    cy.get('[cy="edit-name"]').type("edited name")
    cy.get('[cy="edit-url"]').type("edited url")
    cy.get('[ cy="edit-btn"]').click()
    cy.contains("edited url").should('be.visible')
    cy.contains("edited name").should('be.visible')
    cy.get('[cy="link-delete"]').first().click()
    cy.contains("edited url").should('not.exist')
    cy.contains("edited name").should('not.exist')
    cy.get('[cy="profile"]').click()

    // profile: edit page
    cy.get('[cy="img"]').selectFile("cypress/e2e/testdata/test-img.jpg")
    cy.get('[cy="bio"]').type("hello, it's me")
    cy.get('[cy="profile-update-btn"]').click()
    cy.get('[cy="img"]').should('be.empty')
    cy.get('[cy="bio"]').invoke('val')
    .then(val=>{    
      const myVal = val;      
      expect(myVal).to.equal("hello, it's me");
    })


    // profile: show analysis 
    cy.get('[cy="profile-analysis"]').click()
    cy.contains("Click Count").should('be.visible')
    cy.get('[cy="profile-edit"]').click()
    
    // profile: delete account
    cy.get('[cy="profile-delete-btn"]').click()
  })
})

