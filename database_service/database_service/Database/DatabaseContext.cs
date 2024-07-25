using DatabaseAPI.Database.Models;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Database;

public class DatabaseContext(DbContextOptions<DatabaseContext> options) : DbContext(options)
{
	public DbSet<Achievement> Achievements { get; set; }
	public DbSet<CoinsRequest> CoinsRequests { get; set; }
	public DbSet<Merch> Merches { get; set; }
	public DbSet<MerchRequest> MerchRequests { get; set; }
	public DbSet<Role> Roles { get; set; }
	public DbSet<Status> Statuses { get; set; }
	public DbSet<User> Users { get; set; }
}
