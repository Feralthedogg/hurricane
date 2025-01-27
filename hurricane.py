import numpy as np
import matplotlib.pyplot as plt

num_nodes = 100
radius_step = 0.2
height_step = 0.1
theta_step = np.pi / 10

x_vals, y_vals, z_vals = [], [], []

for i in range(num_nodes):
    r = i * radius_step
    theta = i * theta_step
    z = i * height_step

    x = r * np.cos(theta)
    y = r * np.sin(theta)

    x_vals.append(x)
    y_vals.append(y)
    z_vals.append(z)

fig = plt.figure(figsize=(8, 6))
ax = fig.add_subplot(111, projection='3d')

ax.scatter(x_vals, y_vals, z_vals, color='blue', linewidth=1)

ax.set_xlabel("X (cos 0)")
ax.set_ylabel("Y (sin 0)")
ax.set_zlabel("Z (Height)")
ax.set_title("3D Spiral (Hurricane Structure)")

plt.show()
