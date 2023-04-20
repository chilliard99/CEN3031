describe('visit poker terms', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('visits poker terms tab', () => {
        cy.get("#mat-tab-label-0-2").click()
    });

  })