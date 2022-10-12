import React,{ createContext, useContext, ReactNode, useState } from "react";

type authContextType = {
    user: boolean | null;
    login: (id: number) => void;
    logout: () => void;
    id: number;
};

const authContextDefaultValues: authContextType = {
    user: null,
    login: (id: number) => {},
    logout: () => {},
    id: 0
};

const AuthContext = createContext<authContextType>(authContextDefaultValues);

export function useAuth() {
    return useContext(AuthContext);
}

type Props = {
    children: ReactNode;
};

export default function AuthProvider({ children }: Props) {
    const [user, setUser] = useState<boolean | null>(null);
    const [id, setId] = useState<number>(0);

    const login = (id: number) => {
        setUser(true);
        setId(id);
    };

    const logout = () => {
        setUser(false);
    };

    const value = {
        user,
        login,
        logout,
        id
    };

    return (
        <>
            <AuthContext.Provider value={value}>
                {children}
            </AuthContext.Provider>
        </>
    );
}