describe('test royal', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('Click Suits', () => {
        cy.get("#defaultcard0").click()
        cy.get("#diamond").click()
        cy.get("#10").click()

        cy.get("#defaultcard1").click()
        cy.get("#diamond").click()
        cy.get("#jack").click()

        cy.get("#defaultcard2").click()
        cy.get("#diamond").click()
        cy.get("#queen").click()

        cy.get("#defaultcard3").click()
        cy.get("#diamond").click()
        cy.get("#king").click()

        cy.get("#defaultcard4").click()
        cy.get("#diamond").click()
        cy.get("#ace").click()

        cy.get("#defaultcard5").click()
        cy.get("#diamond").click()
        cy.get("#2").click()

        cy.get("#defaultcard6").click()
        cy.get("#diamond").click()
        cy.get("#3").click()
        cy.get("#mat-tab-label-0-1").click()
        cy.get("#mat-tab-label-0-0").click()
        cy.get("#removeAll").click()
    });

  })