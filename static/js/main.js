async function reboot() {
  const response = await fetch('/api/reboot', { method: 'POST' });

  if (response.ok) {
    // alert('Your Raspberry Pi will reboot in 60s... Please wait for the Raspberry Pi to restart.');
    window.location.href = '/success.html';
  } else {
    alert('Failed to reboot Raspberry Pi.');
  }
}