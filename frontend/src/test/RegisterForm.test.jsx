import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import RegisterForm from "./RegisterForm";

// Mock de console.log pour vérifier l'appel
jest.spyOn(console, "log").mockImplementation(() => {});

describe("RegisterForm", () => {
  test("soumet les bonnes données", () => {
    render(<RegisterForm />);

    const emailInput = screen.getByLabelText(/adresse email/i);
    const usernameInput = screen.getByLabelText(/nom d'utilisateur/i);
    const passwordInput = screen.getByLabelText(/mot de passe/i);
    const submitButton = screen.getByRole("button", { name: /s'inscrire/i });

    // Simule l'interaction utilisateur
    fireEvent.change(emailInput, { target: { value: "test@mail.com" } });
    fireEvent.change(usernameInput, { target: { value: "testuser" } });
    fireEvent.change(passwordInput, { target: { value: "password123" } });
    fireEvent.click(submitButton);

    // Vérifie que console.log est appelé avec les bonnes données
    expect(console.log).toHaveBeenCalledWith("Tentative d'inscription :", {
      email: "test@mail.com",
      username: "testuser",
      password: "password123",
    });
  });
});
