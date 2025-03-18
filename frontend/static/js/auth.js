// checkAuthStatus validates the current user's authentication status.
export async function checkAuthStatus() {
  try {
    // Retrieve session token from session storage
    const sessionToken = sessionStorage.getItem('session_token');

    const headers = {
      'X-Requested-With': 'XMLHttpRequest',
    };

    if (sessionToken) {
      headers.Authorization = `Bearer ${sessionToken}`;
    }

    const response = await fetch('/api/auth-status', {
      headers: headers,
    });

    if (response.ok) {
      const data = await response.json();
      return {
        authenticated: data.authenticated,
        userId: data.userId || null,
        message: data.message || '',
      };
    } else {
      return {
        authenticated: false,
        userId: null,
        message: 'Failed to check authentication status',
      };
    }
  } catch (error) {
    console.error('Auth status check error:', error);
    return {
      authenticated: false,
      userId: null,
      message: 'Error checking authentication status',
    };
  }
}

// Logs out the current user.
export async function logout() {
  try {
    // Retrieve session token from session storage
    const sessionToken = sessionStorage.getItem('session_token');

    const headers = {
      'Content-Type': 'application/json',
      'X-Requested-With': 'XMLHttpRequest',
    };

    if (sessionToken) {
      headers.Authorization = `Bearer ${sessionToken}`;
    }

    const response = await fetch('/api/auth-status', {
      headers: headers,
    });

    // Clear session storage on logout
    sessionStorage.removeItem('session_token');

    const data = await response.json();
    return {
      success: data.success,
      message: data.message || 'Logged out successfully',
    };
  } catch (error) {
    console.error('Logout error:', error);
    return {
      success: false,
      message: 'Error during logout',
    };
  }
}

// Redirects unauthenticated users when attempting to access protected routes.
export async function requireAuth(redirectPath = '/sign-in') {
  const { authenticated } = await checkAuthStatus();

  if (!authenticated) {
    // If not authenticated, redirect to the specified path
    history.pushState(null, null, redirectPath);
    return false;
  }

  return true;
}
