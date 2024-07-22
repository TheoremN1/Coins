using Microsoft.EntityFrameworkCore;
using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.EntityFrameworkCore.Storage;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Newtonsoft.Json;

namespace DatabaseAPI;

public class Program
{
	public static void Main(string[] args)
	{
		var builder = WebApplication.CreateBuilder(args);
		using (var streamReader = new StreamReader(Path.Join("Configs", "database.json")))
		{
			var json = streamReader.ReadToEnd();
			var items = JsonConvert.DeserializeObject<Dictionary<string, string>>(json)
				?? throw new NullReferenceException("Что-то не так с файлом database.json");
			var stringConnection =
				$"Host={items["host"]};" +
				$"Port={items["port"]};" +
				$"Database={items["database"]};" +
				$"Username={items["username"]};" +
				$"Password={items["password"]}";
			builder.Services.AddSingleton(_ => {
				var options = new DbContextOptionsBuilder<DatabaseContext>();
				options.UseNpgsql(stringConnection);
				return new DatabaseContext(options.Options);
			});
		}
		builder.Services.AddControllers();

		var app = builder.Build();

		var context = app.Services.GetRequiredService<DatabaseContext>();
		var databaseCreator = (RelationalDatabaseCreator)context.Database.GetService<IDatabaseCreator>();
		// TODO: \/ При релизе убрать \/
		databaseCreator.EnsureDeleted(); 
		databaseCreator.EnsureCreated();
		//		 /\ При релизе убрать /\
		using (var streamReader = new StreamReader(Path.Join("Configs", "roles.json")))
		{
			var json = streamReader.ReadToEnd();
			var items = JsonConvert.DeserializeObject<List<Dictionary<string, string>>>(json)
				?? throw new NullReferenceException("Что-то не так с файлом roles.json");
			foreach(var item in items)
			{
				if (item is null)
					throw new NullReferenceException("Что-то не так с файлом roles.json");
				context.Roles.Add(new Role(item["key"], item["name"]));
			}
		};
        using (var streamReader = new StreamReader(Path.Join("Configs", "statuses.json")))
        {
            var json = streamReader.ReadToEnd();
            var items = JsonConvert.DeserializeObject<List<Dictionary<string, string>>>(json)
                ?? throw new NullReferenceException("Что-то не так с файлом statuses.json");
            foreach (var item in items)
            {
                if (item is null)
                    throw new NullReferenceException("Что-то не так с файлом statuses.json");
                context.Statuses.Add(new Status(item["key"], item["name"]));
            }
        };
		context.SaveChanges();

		app.UseAuthorization();
		app.MapControllers();
		app.Run();
	}
}
