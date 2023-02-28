describe('Add Card', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('Click Suits', () => {
        cy.get('#SuitSelect').click()
        cy.get('#diamond').click()
        cy.get('#ValsSelect').click()
        cy.get('#ace').click()
        cy.get('#add_card').click()
    });

  })